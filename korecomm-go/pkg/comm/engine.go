package comm

import (
	"fmt"
	"github.com/hegemone/kore-poc/korecomm-go/pkg/config"
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

// Engine is the heart of korecomm. It's responsible for routing traffic amongst
// buffers in a concurrent way, as well as the loading and execution of extensions.
type Engine struct {
	// Messaging buffers
	rawIngressBuffer chan rawIngressBufferMsg
	ingressBuffer    chan ingressBufferMsg
	egressBuffer     chan egressBufferMsg

	// Extensions
	plugins  map[string]*Plugin
	adapters map[string]*Adapter
}

// NewEngine creates a new Engine.
func NewEngine() *Engine {
	// Configurable size of the internal message buffers
	bufferSize := config.GetEngineConfig().BufferSize

	return &Engine{
		rawIngressBuffer: make(chan rawIngressBufferMsg, bufferSize),
		ingressBuffer:    make(chan ingressBufferMsg, bufferSize),
		egressBuffer:     make(chan egressBufferMsg, bufferSize),
		plugins:          make(map[string]*Plugin),
		adapters:         make(map[string]*Adapter),
	}
}

// LoadExtensions will attempt to load enabled plugins and extensions. Includes
// extension init (used for things like establishing connections with platforms).
func (e *Engine) LoadExtensions() error {
	log.Info("Loading extensions")
	if err := e.loadPlugins(); err != nil {
		return err
	}
	if err := e.loadAdapters(); err != nil {
		return err
	}
	return nil
}

// Start will cause the engine to start listening on all successfully loaded
// adapters. On the receipt of any new message from an adapter, it will parse
// the message and determine if the contents are a command. If the message does
// contain a command, it will be transformed to an `IngressMessage` and routed
// to matching plugin commands. If the plugin sends back a message to the
// originator, it will be transformed to an `EgressMessage` and routed to the
// originating adapter for transmission via the client.
func (e *Engine) Start() {
	log.Debug("Engine::Start")

	// Spawn listening routines for each adapter
	for _, adapter := range e.adapters {
		adapterCh := make(chan RawIngressMessage)

		go func(adapter *Adapter, adapterCh chan RawIngressMessage) {
			// Tell the adapter to start listening and sending messages back via
			// their own ingress channel. Listen should be non-blocking!
			adapter.Listen(adapterCh)

			// Engine listens to the N channels the adapters are transmitting on
			// for RawIngressMessages. Adapter channels are fanned-in to the
			// rawIngressBuffer for parsing.
			for rim := range adapterCh {
				e.rawIngressBuffer <- rawIngressBufferMsg{adapter.Name, rim}
			}
		}(adapter, adapterCh)
	}

	// Wire up messaging events to their handlers
	for {
		select {
		case m := <-e.rawIngressBuffer:
			e.handleRawIngress(m)
		case m := <-e.ingressBuffer:
			e.handleIngress(m)
		case m := <-e.egressBuffer:
			e.handleEgress(m)
		}
	}
}

// Buffer messages are internal messaging types usually containing a public
// payload + some kind of metadata, ex: to facilitate routing

type rawIngressBufferMsg struct {
	AdapterName       string
	RawIngressMessage RawIngressMessage
}

type ingressBufferMsg struct {
	// NOTE: It's possible in the future we'll want some additional metadata
	// on this to assist the engine in routing a cmd to a plugin. Right now,
	// it's just the cmd less the trigger prefix, which get's matched on the
	// plugin's `CmdManifest`
	IngressMessage IngressMessage
}

type egressBufferMsg struct {
	Originator    Originator
	EgressMessage EgressMessage
}

// handleRawIngress main function is to filter commands from raw messages.
// If a message is determined to be a command, it is parsed and structured as
// an `IngressMessage`, then passed to the ingressBuffer for further handling.
func (e *Engine) handleRawIngress(m rawIngressBufferMsg) {
	go func() {
		adapterName := m.AdapterName
		rm := m.RawIngressMessage

		if !isCmd(rm.RawContent) {
			return
		}

		if string(rm.RawContent[0]) != adapterCmdTriggerPrefix {
			log.Warningf(
				"raw content was flagged as a command, but does not contain trigger prefix, skipping...",
			)
			log.Warning(rm.RawContent)
			return
		}

		content := parseRawContent(rm.RawContent)

		e.ingressBuffer <- ingressBufferMsg{
			IngressMessage: IngressMessage{
				Originator: Originator{Identity: rm.Identity, AdapterName: adapterName},
				Content:    content,
			},
		}
	}()
}

func parseRawContent(rawContent string) string {
	// NOTE: It's possible in the future we'll want more processing of the raw content
	// some kind of metadata that might be useful for the engine to route the cmd
	// to the plugin. for now
	return rawContent[1:len(rawContent)]
}

// handleIngress is responsible for routing `IngressMessage`s to the set of
// matching plugin cmds that have been registered. If the `CmdDelegate` passed
// to the plugin cmd contains a response, it will construct an `EgressMessage`
// and push to the `Engine`'s egress buffer for dispatch to the relevant adapter.
func (e *Engine) handleIngress(ibm ingressBufferMsg) {
	im := ibm.IngressMessage
	log.Debugf("Engine::handleIngress: %+v", im)

	go func() {
		cmdMatches := e.applyCmdManifests(im.Content)

		for _, cmdMatch := range cmdMatches {
			delegate := NewCmdDelegate(im, cmdMatch.Submatches)

			// Execute plugin command and pass delegate as an intermediary
			cmdMatch.CmdFn(&delegate)

			// If the plugin has sent a response to the delegate, let's build
			// an `EgressMessage` and push that onto the outgoing buffer for dispatch
			if delegate.response != "" {
				e.egressBuffer <- egressBufferMsg{
					Originator:    im.Originator,
					EgressMessage: EgressMessage{delegate.response},
				}
			}
		}
	}()
}

type cmdMatch struct {
	CmdFn      CmdFn
	Submatches []string
}

// applyCmdManifests runs the content against all registered plugin `CmdLink`s
// to determine the set of plugin cmd's that need to be executed.
func (e *Engine) applyCmdManifests(content string) []cmdMatch {
	matches := make([]cmdMatch, 0)

	for _, plugin := range e.plugins {
		for _ /*cmdName*/, cmdLink := range plugin.CmdManifest {
			re := cmdLink.Regexp
			subm := re.FindStringSubmatch(content)

			if len(subm) > 0 {
				matches = append(matches, cmdMatch{
					CmdFn:      cmdLink.CmdFn,
					Submatches: subm,
				})
			}
		}
	}

	return matches
}

// handleEgress simply routes an `EgressMessage` off the egressBuffer to an
// adapter for transmission.
func (e *Engine) handleEgress(ebm egressBufferMsg) {
	log.Debugf("Engine::handleEgress: %+v", ebm)
	go func() {
		e.adapters[ebm.Originator.AdapterName].SendMessage(ebm.EgressMessage)
	}()
}

// TODO: load{Plugins,Adapters} are almost identical. Should make extension
// loading generic.
func (e *Engine) loadPlugins() error {
	config := config.GetPluginConfig()
	log.Infof("Loading plugins from: %v", config.Dir)

	// TODO: Check that requested plugins are available in dir, log if not
	for _, pluginName := range config.Enabled {
		log.Infof("-> %v", pluginName)
		pluginFile := filepath.Join(
			config.Dir,
			fmt.Sprintf("%s.so", pluginName),
		)

		loadedPlugin, err := LoadPlugin(pluginFile)
		if err != nil {
			// TODO: Probably want this to be more resilient so the comm server can
			// skip problematic plugins while still loading valid ones.
			return err
		}

		e.plugins[loadedPlugin.Name] = loadedPlugin
	}

	log.Info("Successfully loaded plugins:")
	for pluginName, _ := range e.plugins {
		log.Infof("-> %s", pluginName)
	}

	return nil
}

func (e *Engine) loadAdapters() error {
	config := config.GetAdapterConfig()
	log.Infof("Loading adapters from: %v", config.Dir)

	// TODO: Check that requested adapters are available in dir, log if not
	for _, adapterName := range config.Enabled {
		log.Infof("-> %v", adapterName)
		adapterFile := filepath.Join(
			config.Dir,
			fmt.Sprintf("%s.so", adapterName),
		)
		log.Infof("file: %s", adapterFile)

		loadedAdapter, err := LoadAdapter(adapterFile)
		if err != nil {
			// TODO: Probably want this to be more resilient so the comm server can
			// skip problematic adapters while still loading valid ones.
			return err
		}

		// Initialize the adapters in determinic sequence. Init should be used for
		// things like establishing a connection with an external platform.
		loadedAdapter.Init()
		e.adapters[loadedAdapter.Name] = loadedAdapter
	}

	log.Info("Successfully loaded adapters:")
	for adapterName, _ := range e.adapters {
		log.Infof("-> %s", adapterName)
	}
	return nil
}

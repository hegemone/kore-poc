package comm

import (
	"fmt"
	goplugin "plugin"
	"regexp"
)

// The prefix trigger used to denote a cmd.
// Example: !droplet start skynet
const adapterCmdTriggerPrefix = "!"

// Regexp applied to check isCmd.
var adapterCmdRegexp, _ = regexp.Compile(fmt.Sprintf("^%s\\S*($| )",
	adapterCmdTriggerPrefix))

// Adapter is an abstraction that should be implemented to present a standard
// interface to the comm server for communicating to, and from external platforms.
// Similar to the `Plugin`, it is a facade type that delegates  actions like
// sending and receiving messages to concrete implementations dynamically loaded
// from shared libraries.
type Adapter struct {
	Name string

	fnInit        func()
	fnSendMessage func(string)
	fnName        func() string
	fnListen      func(chan<- RawIngressMessage)
}

// Listen is the public trigger that initiates an adapter to start listening
// to external platform events. It should be implemented as non-blocking and
// push `RawIngressMessage`s to the inChan on the receipt of raw messages
// from the external platform.
func (a *Adapter) Listen(inChan chan<- RawIngressMessage) {
	// Possibly some common logic an Adapter might want to do instead of having
	// the engine call the raw plugin listen directly
	// NOTE: Engine has already handled spawning the listen routines in their
	// own goroutines, so they're running concurrently and/or in parallel.
	// TODO: Engine probably needs to handle the case of adapters being poorly
	// written and immediately having their channels close or leaked. fnListen
	// is expected to be long lived.
	a.fnListen(inChan)
}

// SendMessage is the public trigger indicating a dynamically loaded adapter
// should transmit an `EgressMessage` to its platform. Dynamically loaded
// adapters must define how that is done.
func (a *Adapter) SendMessage(emsg EgressMessage) {
	// TODO: General processing
	a.fnSendMessage(emsg.Serialize())
}

// LoadAdapter loads adapter behavior from a given .so adapter file
func LoadAdapter(adapterFile string) (*Adapter, error) {
	// TODO: Need a *lot* of validation here to make sure a bad adapter doesn't
	// just crash the server.
	// -> Actually confirm the casts are valid and these functions look like they should?
	// TODO: Can the hardcoded pattern of $PROPERTY Lookup -> Cast be made more elegant?
	a := Adapter{}

	rawGoPlugin, err := goplugin.Open(adapterFile)
	if err != nil {
		return nil, err
	}

	nameSym, err := rawGoPlugin.Lookup("Name")
	if err != nil {
		return nil, err
	}
	a.fnName = nameSym.(func() string)
	a.Name = a.fnName()

	listenSym, err := rawGoPlugin.Lookup("Listen")
	if err != nil {
		return nil, err
	}
	a.fnListen = listenSym.(func(chan<- RawIngressMessage))

	sendMessageSym, err := rawGoPlugin.Lookup("SendMessage")
	if err != nil {
		return nil, err
	}
	a.fnSendMessage = sendMessageSym.(func(string))

	initSym, err := rawGoPlugin.Lookup("Init")
	if err != nil {
		return nil, err
	}
	a.fnInit = initSym.(func())

	return &a, nil
}

// Init gives adapters an opportunity to implement an initialization hook.
// Typically used for things like setting up initial connection to external
// platform APIs, checking for the presence of credentials and their validity, etc.
func (a *Adapter) Init() {
	a.fnInit()
}

func isCmd(rawContent string) bool {
	// isCmd is where the adapter defines whether or not raw content is indeed a Cmd.
	return adapterCmdRegexp.MatchString(rawContent)
}

package comm

import (
	log "github.com/sirupsen/logrus"
	goplugin "plugin"
	"regexp"
)

// CmdFn is the required function signature of a plugin command. It should accept
// a `*CmdDelegate` for interfacing back with the `Engine`.
type CmdFn func(*CmdDelegate)

// CmdLink pairs a regexp to a CmdFn. If the engine matches the regexp against
// an `IngressMessage`, it will route the command to the `CmdLink`'s `CmdFn`.
type CmdLink struct {
	Regexp *regexp.Regexp
	CmdFn  CmdFn
}

// Plugin is the primary abstraction representing dynamic, extensible behavioral features.
// In reality, it's a facade that presents a controller public interface for
// consumers, while delegating much of its functionality to dynamically loaded
// functions sourced from a shared library.
type Plugin struct {
	Name        string
	Help        string
	CmdManifest map[string]CmdLink

	fnName        func() string
	fnHelp        func() string
	fnCmdManifest func() map[string]string
}

// LoadPlugin loads dynamic plugin behavior from a given .so plugin file
func LoadPlugin(pluginFile string) (*Plugin, error) {
	// TODO: Need a *lot* of validation here to make sure a bad plugin doesn't
	// just crash the server.
	// -> Actually confirm the casts are valid and these functions look like they should?
	// TODO: Can the hardcoded pattern of $PROPERTY Lookup -> Cast be made more elegant?
	p := Plugin{}

	rawGoPlugin, err := goplugin.Open(pluginFile)
	if err != nil {
		return nil, err
	}

	nameSym, err := rawGoPlugin.Lookup("Name")
	if err != nil {
		return nil, err
	}
	p.fnName = nameSym.(func() string)
	p.Name = p.fnName()

	helpSym, err := rawGoPlugin.Lookup("Help")
	if err != nil {
		return nil, err
	}
	p.fnHelp = helpSym.(func() string)
	p.Help = p.fnHelp()

	cmdManifestSym, err := rawGoPlugin.Lookup("CmdManifest")
	if err != nil {
		return nil, err
	}
	p.fnCmdManifest = cmdManifestSym.(func() map[string]string)

	p.CmdManifest = make(map[string]CmdLink)
	for cmdRegexStr, cmdFnName := range p.fnCmdManifest() {
		cmdFnSym, err := rawGoPlugin.Lookup(cmdFnName)
		if err != nil {
			log.Error("Error occurred while looking up command for plugin %s: %s", p.Name, err.Error())
			continue
		}

		cmdRegex, _ := regexp.Compile(cmdRegexStr)    // TODO: Handle failed regex compilation
		cmdFn := CmdFn(cmdFnSym.(func(*CmdDelegate))) // TODO: Handle failed cast

		// TODO: Error handle more than one command named the same thing
		p.CmdManifest[cmdFnName] = CmdLink{
			Regexp: cmdRegex,
			CmdFn:  cmdFn,
		}
	}

	return &p, nil
}

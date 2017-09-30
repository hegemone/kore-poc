package comm

import (
	log "github.com/sirupsen/logrus"
)

// CmdDelegate is the object passed to plugin cmds that acts as an intermediary
// between the `Engine` and the `Plugin`. It is the primary interface a plugin
// author interfaces with the comm server.
type CmdDelegate struct {
	IngressMessage IngressMessage
	Submatches     []string

	response string
}

// NewCmdDelegate creates a new `CmdDelegate` from an `IngressMessage` and
// a submatch array, produced from the `CmdLink.Regexp` match.
func NewCmdDelegate(im IngressMessage, subm []string) CmdDelegate {
	return CmdDelegate{
		IngressMessage: im,
		Submatches:     subm,
		response:       "",
	}
}

// SendResponse is used by plugin cmds to communicate messages back to the
// originating platform that triggered the command. It accepts a simple
// string response.
func (d *CmdDelegate) SendResponse(response string) {
	log.Debugf("CmdDelegate::SendResponse: %s", response)
	d.response = response
}

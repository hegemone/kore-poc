// Example plugin. Implements the classic bacon cinch plugin.
package main

import (
	"fmt"
	"github.com/hegemone/kore-poc/korecomm-go/pkg/comm"
	log "github.com/sirupsen/logrus"
)

func Name() string {
	return "bacon.plugins.kore.nsk.io"
}

func CmdManifest() map[string]string {
	return map[string]string{
		`bacon$`:        "CmdBacon",
		`bacon\s+(\S+)`: "CmdBaconGift",
	}
}

func Help() string {
	return "Usage: !bacon [user]"
}

func CmdBacon(p *comm.CmdDelegate) {
	log.Infof("bacon.plugins::CmdBacon, IngressMessage: %+v", p.IngressMessage)

	msg := p.IngressMessage
	identity := msg.Originator.Identity

	response := fmt.Sprintf(
		"gives %s a strip of delicious bacon.", identity,
	)

	p.SendResponse(response)
}

func CmdBaconGift(p *comm.CmdDelegate) {
	log.Infof("bacon.plugins::CmdBaconGift, IngressMessage: %+v", p.IngressMessage)

	msg := p.IngressMessage
	identity := msg.Originator.Identity
	toUser := p.Submatches[1]

	response := fmt.Sprintf(
		"gives %s a strip of delicious bacon as a gift from %v",
		toUser, identity,
	)

	p.SendResponse(response)
}

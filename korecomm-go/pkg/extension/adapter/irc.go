// Example irc adapter. Expected to be built as a standalone .so.
package main

import (
	"github.com/hegemone/kore-poc/korecomm-go/pkg/comm"
	"github.com/hegemone/kore-poc/korecomm-go/pkg/mock"
	log "github.com/sirupsen/logrus"
)

// pkg level client reference
var ircAdapterClient *mock.PlatformClient

////////////////////////////////////////////////////////////////////////////////
// Concrete Behavioral Implementations
////////////////////////////////////////////////////////////////////////////////

func Init() {
	log.Info("ex-irc.adapters::Init")
	// NOTE: In a real adapter, this client is probably some kind of API used
	// to actually speak to irc. For the purposes of this POC, we are using
	// a mocked "PlatformClient" that is designed behave like that platform API might.
	// It is driven by Stdin thanks to the StdinDemux.
	ircAdapterClient = mock.NewPlatformClient("irc")
}

func Name() string {
	return "ex-irc.adapters.kore.nsk.io"
}

func Listen(ingressCh chan<- comm.RawIngressMessage) {
	log.Debug("ex-irc.adapters::Listen")

	ircAdapterClient.Connect()

	go func() {
		for clientMsg := range ircAdapterClient.Chat {
			ingressCh <- comm.RawIngressMessage{
				Identity:   clientMsg.User,
				RawContent: clientMsg.Message,
			}
		}
	}()
}

func SendMessage(m string) {
	ircAdapterClient.SendMessage(m)
}

// Example discord adapter. Expected to be built as a standalone .so.
package main

import (
	"github.com/hegemone/kore-poc/korecomm-go/pkg/comm"
	"github.com/hegemone/kore-poc/korecomm-go/pkg/mock"
	log "github.com/sirupsen/logrus"
)

// pkg level client reference
var discordAdapterClient *mock.PlatformClient

////////////////////////////////////////////////////////////////////////////////
// Concrete Behavioral Implementations
////////////////////////////////////////////////////////////////////////////////

func Init() {
	log.Info("ex-discord.adapters::Init")
	// NOTE: In a real adapter, this client is probably some kind of API used
	// to actually speak to discord. For the purposes of this POC, we are using
	// a mocked "PlatformClient" that is designed behave like that platform API might.
	// It is driven by Stdin thanks to the StdinDemux.
	discordAdapterClient = mock.NewPlatformClient("discord")
}

func Name() string {
	return "ex-discord.adapters.kore.nsk.io"
}

func Listen(ingressCh chan<- comm.RawIngressMessage) {
	log.Debug("ex-discord.adapters::Listen")

	discordAdapterClient.Connect()

	go func() {
		for clientMsg := range discordAdapterClient.Chat {
			ingressCh <- comm.RawIngressMessage{
				Identity:   clientMsg.User,
				RawContent: clientMsg.Message,
			}
		}
	}()
}

func SendMessage(m string) {
	discordAdapterClient.SendMessage(m)
}

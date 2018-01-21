// Example plugin. Implements the classic bacon cinch plugin.
package main

import (
	"context"
	"fmt"
	"github.com/hegemone/kore-poc/korecomm-go/pkg/comm"
	ioclient "github.com/hegemone/kore-poc/koredata-goa/client"
	log "github.com/sirupsen/logrus"
	"strings"
)

func Name() string {
	return "quotes.plugins.kore.firstofth300.io"
}

func CmdManifest() map[string]string {
	return map[string]string{
		`quote\s+(\S+)`: "CmdGetQuote",
	}
}

func Help() string {
	return "Usage: !quote [quotee]"
}

func CmdGetQuote(p *comm.CmdDelegate) {
	log.Infof("quotes.plugins::CmdGetQuote, IngressMessage: %+v", p.IngressMessage)

	quotee := strings.Fields(p.Submatches[0])[1]

	ctx := context.TODO()
	koreClient := ioclient.New(nil)
	koreClient.Client.Host = "127.0.0.1:8080"

	response, err := koreClient.ListByIDQuote(ctx, quotee)

	if err != nil {
		log.Infof("error on calling data server: %s", err)
	}

	if response != nil {
		quote, err := koreClient.DecodeQuote(response)

		if err != nil {
			log.Infof("unable to decode quote: %s", err)
		}
		response := fmt.Sprintf(
			"%s - %s",
			*quote.Quote, *quote.Name,
		)

		p.SendResponse(response)
	}
}

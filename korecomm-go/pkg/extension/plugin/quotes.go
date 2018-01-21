// Example plugin. Implements the classic bacon cinch plugin.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/hegemone/kore-poc/korecomm-go/pkg/comm"
	"github.com/hegemone/kore-poc/koredata-goa/app"
	log "github.com/sirupsen/logrus"
	"net/http"
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

	koreClient := http.Client{}

	quotesJson, err := koreClient.Get("http://localhost:8080/quotes/" + quotee)

	if err != nil {
		log.Infof("unable to fetch quotes json: %s", err)
	}

	if quotesJson != nil {
		var quotes app.JSON
		log.Info("Decoding JSON")
		decoder := json.NewDecoder(quotesJson.Body)
		decoder.Decode(&quotes)
		log.Info("Decoded JSON")
		log.Info(len(quotes.Quotes))
		response := fmt.Sprintf(
			"%s - %s",
			*quotes.Quotes[0].Quote, *quotes.Quotes[0].Name,
		)

		p.SendResponse(response)
	}
}

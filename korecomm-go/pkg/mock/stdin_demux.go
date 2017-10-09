package mock

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"sync"
)

var _stdinDemuxInstance *StdinDemux
var _stdinDemuxOnce sync.Once

type regReq struct {
	name   string
	chatCh chan<- ChatMessage
}

type StdinDemux struct {
	clients  map[string]chan<- ChatMessage
	regReqCh chan regReq
}

func StdinDemuxInstance() *StdinDemux {
	// Threadsafe lazy accessor
	_stdinDemuxOnce.Do(func() {
		_stdinDemuxInstance = &StdinDemux{
			clients:  make(map[string]chan<- ChatMessage),
			regReqCh: make(chan regReq),
		}
	})
	return _stdinDemuxInstance
}

func (d *StdinDemux) Listen() {
	go func() {
		log.Debug("StdinDemux::Listen")

		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			clientName, chatMsg, err := structuredMsg(text)
			if err != nil {
				log.Error(err.Error())
				continue
			}

			chatCh, ok := d.clients[clientName]
			if !ok {
				log.Infof("Client %s not registered with stdin demux, skipping...", clientName)
				continue
			}

			log.Infof("Demux -> %s", clientName)
			chatCh <- *chatMsg
		}
	}()

	// Listen for reg requests, which commonly come in concurrently, so need to
	// be orchestrated by a chan
	go func() {
		log.Info("Started the reg req func...")
		for r := range d.regReqCh {
			log.Debugf("Got regReq: %+v", r)
			d.clients[r.name] = r.chatCh
		}
	}()
}

func structuredMsg(text string) (string, *ChatMessage, error) {
	split := strings.Split(text, " ")
	if !(len(split) > 1) {
		return "", nil, errors.New("Must send stdin message in format of '<adapter_name> <content>'")
	}

	clientName := split[0]
	strippedSplit := split[1:len(split)]
	message := strings.Join(strippedSplit, " ")

	return clientName, &ChatMessage{
		User:    fmt.Sprintf("%s-user", clientName),
		Message: message,
	}, nil
}

func (d *StdinDemux) Register(name string, chatCh chan<- ChatMessage) {
	log.Debugf("StdinDemux::Register, name: %s", name)
	d.regReqCh <- regReq{name, chatCh}
}

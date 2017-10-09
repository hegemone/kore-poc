package mock

import (
	log "github.com/sirupsen/logrus"
)

// NOTE: This is a mock client that mocks out a vendored API that serves as
// a client into a particular platform, say Discord or IRC. Normally each
// platform's client would look completely different, but this is an Stdin
// based client that demos dynamic, incomming content.

type ChatMessage struct {
	User    string
	Message string
}

type PlatformClient struct {
	Chat chan ChatMessage

	name  string
	demux *StdinDemux
}

func NewPlatformClient(name string) *PlatformClient {
	return &PlatformClient{
		name: name,
		Chat: make(chan ChatMessage),
	}
}

func (c *PlatformClient) Connect() {
	log.Debugf("%s::PlatformClient::Connect", c.name)
	c.demux = StdinDemuxInstance()
	c.demux.Register(c.name, c.Chat)
}

func (c *PlatformClient) SendMessage(m string) {
	log.Infof("%s client got a message from the adapter! [ %s ]", c.name, m)
}

package comm

// Originator contains originating information for an `IngressMessage`. It
// contains the "identity" of the person/system that triggered the incoming
// message, in addition to the adapter name that procuded the event.
type Originator struct {
	Identity    string
	AdapterName string
}

// IngressMessage is the structured, parsed message representing an incoming
// command to be processed. An `IngressMessage` has been determined by the
// system to have content containing a command.
type IngressMessage struct {
	Content    string
	Originator Originator
}

// RawIngressMessage - Raw, unprocessed message passed from the adapter to the
// engine. Has not yet been parsed to determine if the message is a cmd or not.
type RawIngressMessage struct {
	Identity   string
	RawContent string
}

package comm

// EgressMessage is the structured outgoing message Adapters should implement
// how to handle in their `SendMessage` function.
type EgressMessage struct {
	Content string
}

// Serialize simply serializes an `EgressMessage`.
func (e *EgressMessage) Serialize() string {
	// NOTE: Might want to expand on this in the future
	return e.Content
}

package midiv1

// MessageBuilder represents MIDI message data that can be both marshalled and unmarshalled.
type MessageBuilder interface {
	MessageMarshaler
	MessageUnmarshaler
}

// RunningStatusMessageBuilder represents running status MIDI message data that can be both marshalled and unmarshalled.
type RunningStatusMessageBuilder interface {
	RunningStatusMessageMarshaler
	RunningStatusMessageUnmarshaler
}

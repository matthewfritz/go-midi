package midiv1

import "errors"

var (
	// ErrMarshallingMessage represents an error marshalling a MIDI message.
	ErrMarshallingMessage error = errors.New("error marshalling MIDI message")
)

// MessageMarshaler represents MIDI message data that can be marshalled.
type MessageMarshaler interface {
	// MarshalMIDI marshalls a MIDI message into its raw bytes.
	MarshalMIDI() ([]byte, error)
}

// RunningStatusMessageMarshaler represents running status MIDI message data that can be marshalled.
type RunningStatusMessageMarshaler interface {
	// MarshalRunningStatusMIDI marshalls a running status MIDI message into its raw bytes.
	MarshalRunningStatusMIDI() ([]byte, error)
}

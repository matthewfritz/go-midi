package midiv1

import "errors"

var (
	// ErrUnmarshallingMessage represents an error unmarshalling a MIDI message.
	ErrUnmarshallingMessage error = errors.New("error unmarshalling MIDI message")
)

// MessageUnmarshaler represents MIDI message data that can be unmarshalled.
type MessageUnmarshaler interface {
	// UnmarshalMIDI unmarshalls raw bytes into a MIDI message.
	UnmarshalMIDI(b []byte) error
}

// RunningStatusMessageUnmarshaler represents running status MIDI message data that can be unmarshalled.
type RunningStatusMessageUnmarshaler interface {
	// UnmarshalRunningStatusMIDI unmarshalls raw bytes into a running status MIDI message.
	UnmarshalRunningStatusMIDI(b []byte) error
}

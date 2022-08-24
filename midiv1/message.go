package midiv1

import (
	"errors"
	"fmt"
)

// MessageBuilder represents MIDI message data that can be both marshalled and unmarshalled.
type MessageBuilder interface {
	// MarshalMIDI marshalls a MIDI message into its raw bytes
	MarshalMIDI() ([]byte, error)

	// UnmarshalMIDI unmarshalls raw bytes into a MIDI message
	UnmarshalMIDI(b []byte) error
}

var (
	// ErrUnmarshallingMessage represents an error unmarshalling a MIDI message.
	ErrUnmarshallingMessage error = errors.New("error unmarshalling MIDI message")
)

// Nibble represents a set of four bits within a byte
type Nibble byte

// MessageType represents the type of MIDI message by setting only the MSB of the byte.
type MessageType Nibble

const (
	// StatusMessageMSB represents a MIDI status message byte with a most-significant bit of 1.
	StatusMessageMSB MessageType = 0b10000000

	// DataMessageMSB represents a MIDI data message byte with a most-significant bit of 0.
	DataMessageMSB MessageType = 0b0
)

// ByteHasDataMSB returns whether the most-significant bit of the supplied byte is a MIDI data bit.
func ByteHasDataMSB(b byte) bool {
	return !ByteHasStatusMSB(b)
}

// ByteHasStatusMSB returns whether the most-significant bit of the supplied byte is a MIDI status bit.
func ByteHasStatusMSB(b byte) bool {
	// https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827
	return (b & byte(StatusMessageMSB)) == byte(StatusMessageMSB)
}

// Status represents the first four bits of the MIDI message status byte (message type and code).
//
// Example: 0b11010000 (Status message for Channel Pressure)
type Status Nibble

// Channel represents second four bits of the MIDI message status byte (the ID of the channel).
//
// Example: 0b00001101 (index 13 is channel 14)
type Channel Nibble

var (
	// ErrInvalidChannel represents an invalid MIDI channel ID.
	ErrInvalidChannel error = errors.New("invalid MIDI channel")
)

const (
	// MinChannel is the lowest MIDI channel available (index 0 is channel 1).
	MinChannel Channel = 0b0

	// MaxChannel is the highest MIDI channel available (index 15 is channel 16).
	MaxChannel Channel = 0b1111
)

// NewChannel returns a Channel based on the integer argument.
func NewChannel(channel int) (Channel, error) {
	if channel < int(MinChannel) || channel > int(MaxChannel) {
		return MinChannel, fmt.Errorf("valid channels are between %d and %d, inclusive: %w", MinChannel, MaxChannel, ErrInvalidChannel)
	}
	return Channel(channel), nil
}

// NewChannelFromByte returns a Channel based on the byte argument.
func NewChannelFromByte(channel byte) (Channel, error) {
	if channel < byte(MinChannel) || channel > byte(MaxChannel) {
		return MinChannel, fmt.Errorf("valid channels are between %d and %d, inclusive: %w", MinChannel, MaxChannel, ErrInvalidChannel)
	}
	return Channel(channel), nil
}

// ParseChannelFromStatusByte returns a Channel by AND-ing the channel nibble in the status byte with its maximum value.
func ParseChannelFromStatusByte(status byte) (Channel, error) {
	// https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827
	channel, err := NewChannelFromByte(status & byte(MaxChannel))
	if err != nil {
		return MinChannel, fmt.Errorf("invalid channel (%v) from status byte: %w", err, ErrUnmarshallingMessage)
	}
	return channel, nil
}

// Note represents the note number of the MIDI message.
type Note int

var (
	// ErrInvalidNote represents an invalid MIDI note number.
	ErrInvalidNote error = errors.New("invalid MIDI note")
)

const (
	// MinNote is the lowest MIDI note number available.
	MinNote Note = 0

	// MaxNote is the highest MIDI note number available.
	MaxNote Note = 127
)

// NewNote returns a Note based on the integer argument.
func NewNote(note int) (Note, error) {
	if note < int(MinNote) || note > int(MaxNote) {
		return MinNote, fmt.Errorf("valid note numbers are between %d and %d, inclusive: %w", MinNote, MaxNote, ErrInvalidNote)
	}
	return Note(note), nil
}

// NewNoteFromByte returns a Note based on the byte argument.
func NewNoteFromByte(note byte) (Note, error) {
	if note < byte(MinNote) || note > byte(MaxNote) {
		return MinNote, fmt.Errorf("valid note numbers are between %d and %d, inclusive: %w", MinNote, MaxNote, ErrInvalidNote)
	}
	return Note(note), nil
}

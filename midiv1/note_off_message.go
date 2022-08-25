package midiv1

import "fmt"

const (
	// NoteOffMessageStatusCode represents the message code within the status nibble.
	//
	// This is functionally unnecessary because it's a 0x00 byte and will not affect the calculated result
	// when a bitwise OR is performed. It is retained for consistency among other Channel Voice messages.
	NoteOffMessageCode Nibble = 0b00000000

	// NoteOffMessageLength represents the number of bytes in a full Note-Off message.
	NoteOffMessageLength int = 3

	// NoteOffMessageStatusNibble represents the status nibble within the status byte
	NoteOffMessageStatusNibble Status = Status(StatusMessageMSB) | Status(NoteOffMessageCode)
)

// NoteOffMessage represents a Note-On Channel Voice message.
type NoteOffMessage struct {
	// Channel represents the channel number where this message will be sent.
	Channel Channel

	// Note represents the note number that will be sent with this message.
	Note Note

	// Velocity represents the relative strength of the note in this message.
	Velocity Velocity
}

// MarshalMIDI marshalls a NoteOffMessage MIDI message into its raw bytes
func (nom NoteOffMessage) MarshalMIDI() ([]byte, error) {
	return []byte{
		MakeStatusByte(NoteOffMessageStatusNibble, nom.Channel),
		byte(nom.Note),
		byte(nom.Velocity),
	}, nil
}

// MarshalRunningStatusMIDI marshalls a running status MIDI message into its raw bytes.
func (nom NoteOffMessage) MarshalRunningStatusMIDI() ([]byte, error) {
	return []byte{
		byte(nom.Note),
		byte(nom.Velocity),
	}, nil
}

// UnmarshalMIDI unmarshalls raw bytes into a NoteOffMessage struct pointer. Note-Off messages are
// represented by three bytes (left to right): status/channel, note number, note velocity.
//
// Example: []byte{0b10000001, 0b01000000, 0b00100000}
//
// The example forms a Note-Off message for channel 2 (index 1), note number 64, velocity value 32.
func (nom *NoteOffMessage) UnmarshalMIDI(b []byte) error {
	// check the number of bytes in the message
	if len(b) != NoteOffMessageLength {
		return fmt.Errorf("note-off messages are made up of %d bytes, received %d byte(s): %w", NoteOffMessageLength, len(b), ErrUnmarshallingMessage)
	}

	// make sure this is a status byte with the proper MSB
	if !ByteHasStatusMSB(b[0]) {
		return fmt.Errorf("note-off messages must have a status MSB: %w", ErrUnmarshallingMessage)
	}

	// retrieve the channel nibble of the status byte to form the Channel value
	channel, err := ParseChannelFromStatusByte(b[0])
	if err != nil {
		return err
	}

	// form the note number
	note, err := NewNoteFromByte(b[1])
	if err != nil {
		return fmt.Errorf("invalid note number (%v) from note byte: %w", err, ErrUnmarshallingMessage)
	}

	// form the velocity
	vel := NewVelocityFromByte(b[2])

	*nom = NoteOffMessage{
		Channel:  channel,
		Note:     note,
		Velocity: vel,
	}
	return nil
}

// UnmarshalRunningStatusMIDI unmarshalls raw bytes into a NoteOffMessage struct pointer. Note-Off running status messages are
// represented by two bytes (left to right): note number, note velocity.
//
// Example: []byte{0b01000000, 0b00100000}
//
// The example forms a Note-Off running status message for note number 64, velocity value 32.
func (nom *NoteOffMessage) UnmarshalRunningStatusMIDI(b []byte) error {
	// check the number of bytes in the running status message
	if len(b) != NoteOffMessageLength-1 {
		return fmt.Errorf("note-off running status messages are made up of %d bytes, received %d byte(s): %w", NoteOffMessageLength-1, len(b), ErrUnmarshallingMessage)
	}

	// form the note number
	note, err := NewNoteFromByte(b[0])
	if err != nil {
		return fmt.Errorf("invalid note number %#v (%v) from running status note byte: %w", b[0], err, ErrUnmarshallingMessage)
	}

	// form the velocity
	vel := NewVelocityFromByte(b[1])

	*nom = NoteOffMessage{
		Note:     note,
		Velocity: vel,
	}
	return nil
}

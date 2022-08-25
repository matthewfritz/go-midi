package midiv1

import "fmt"

const (
	// NoteOnMessageStatusCode represents the message code within the status nibble
	NoteOnMessageCode Nibble = 0b00010000

	// NoteOnMessageLength represents the number of bytes in a full Note-On message.
	NoteOnMessageLength int = 3

	// NoteOnMessageStatusNibble represents the status nibble within the status byte
	NoteOnMessageStatusNibble Status = Status(StatusMessageMSB) | Status(NoteOnMessageCode)
)

// NoteOnMessage represents a Note-On Channel Voice message.
type NoteOnMessage struct {
	// Channel represents the channel number where this message will be sent.
	Channel Channel

	// Note represents the note number that will be sent with this message.
	Note Note

	// Velocity represents the relative strength of the note in this message.
	Velocity Velocity
}

// MarshalMIDI marshalls a NoteOnMessage MIDI message into its raw bytes
func (nom NoteOnMessage) MarshalMIDI() ([]byte, error) {
	return []byte{
		MakeStatusByte(NoteOnMessageStatusNibble, nom.Channel),
		byte(nom.Note),
		byte(nom.Velocity),
	}, nil
}

// MarshalRunningStatusMIDI marshalls a running status MIDI message into its raw bytes.
func (nom NoteOnMessage) MarshalRunningStatusMIDI() ([]byte, error) {
	return []byte{
		byte(nom.Note),
		byte(nom.Velocity),
	}, nil
}

// UnmarshalMIDI unmarshalls raw bytes into a NoteOnMessage struct pointer. Note-On messages are
// represented by three bytes (left to right): status/channel, note number, note velocity.
//
// Example: []byte{0b10010001, 0b01000000, 0b00100000}
//
// The example forms a Note-On message for channel 2 (index 1), note number 64, velocity value 32.
func (nom *NoteOnMessage) UnmarshalMIDI(b []byte) error {
	// check the number of bytes in the message
	if len(b) != NoteOnMessageLength {
		return fmt.Errorf("note-on messages are made up of %d bytes, received %d byte(s): %w", NoteOnMessageLength, len(b), ErrUnmarshallingMessage)
	}

	// make sure this is a status byte with the proper MSB
	if !ByteHasStatusMSB(b[0]) {
		return fmt.Errorf("note-on messages must have a status MSB: %w", ErrUnmarshallingMessage)
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

	*nom = NoteOnMessage{
		Channel:  channel,
		Note:     note,
		Velocity: vel,
	}
	return nil
}

// UnmarshalRunningStatusMIDI unmarshalls raw bytes into a NoteOnMessage struct pointer. Note-On running status messages are
// represented by two bytes (left to right): note number, note velocity.
//
// Example: []byte{0b01000000, 0b00100000}
//
// The example forms a Note-On running status message for note number 64, velocity value 32.
func (nom *NoteOnMessage) UnmarshalRunningStatusMIDI(b []byte) error {
	// check the number of bytes in the running status message
	if len(b) != NoteOnMessageLength-1 {
		return fmt.Errorf("note-on running status messages are made up of %d bytes, received %d byte(s): %w", NoteOnMessageLength-1, len(b), ErrUnmarshallingMessage)
	}

	// form the note number
	note, err := NewNoteFromByte(b[0])
	if err != nil {
		return fmt.Errorf("invalid note number %#v (%v) from running status note byte: %w", b[0], err, ErrUnmarshallingMessage)
	}

	// form the velocity
	vel := NewVelocityFromByte(b[1])

	*nom = NoteOnMessage{
		Note:     note,
		Velocity: vel,
	}
	return nil
}

package midiv1

import "fmt"

const (
	// NoteOnMessageStatusCode represents the message code within the status nibble
	NoteOnMessageCode Nibble = 0b00010000

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
		byte(NoteOnMessageStatusNibble) | byte(nom.Channel),
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
	// a Note-On message sequence uses three bytes
	if len(b) != 3 {
		return fmt.Errorf("note-on messages are made up of three bytes, received %d byte(s): %w", len(b), ErrUnmarshallingMessage)
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

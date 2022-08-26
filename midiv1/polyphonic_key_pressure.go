package midiv1

import "fmt"

const (
	// PolyphonicKeyPressureMessageStatusCode represents the message code within the status nibble
	PolyphonicKeyPressureMessageCode Nibble = 0b00100000

	// PolyphonicKeyPressureMessageLength represents the number of bytes in a full Polyphonic Key Pressure message.
	PolyphonicKeyPressureMessageLength int = 3

	// PolyphonicKeyPressureMessageStatusNibble represents the status nibble within the status byte
	PolyphonicKeyPressureMessageStatusNibble Status = Status(StatusMessageMSB) | Status(PolyphonicKeyPressureMessageCode)
)

// PolyphonicKeyPressureMessage represents a Note-On Channel Voice message.
type PolyphonicKeyPressureMessage struct {
	// Channel represents the channel number where this message will be sent.
	Channel Channel

	// Note represents the note number that will be sent with this message.
	Note Note

	// Pressure represents the relative applied pressure of the note in this message.
	Pressure Pressure
}

// MarshalMIDI marshalls a PolyphonicKeyPressureMessage MIDI message into its raw bytes
func (pkpm PolyphonicKeyPressureMessage) MarshalMIDI() ([]byte, error) {
	return []byte{
		MakeStatusByte(PolyphonicKeyPressureMessageStatusNibble, pkpm.Channel),
		byte(pkpm.Note),
		byte(pkpm.Pressure),
	}, nil
}

// MarshalRunningStatusMIDI marshalls a running status MIDI message into its raw bytes.
func (pkpm PolyphonicKeyPressureMessage) MarshalRunningStatusMIDI() ([]byte, error) {
	return []byte{
		byte(pkpm.Note),
		byte(pkpm.Pressure),
	}, nil
}

// UnmarshalMIDI unmarshalls raw bytes into a PolyphonicKeyPressureMessage struct pointer. Polyphonic Key Pressure messages are
// represented by three bytes (left to right): status/channel, note number, note pressure.
//
// Example: []byte{0b10100001, 0b01000000, 0b00100000}
//
// The example forms a Polyphonic Key Pressure message for channel 2 (index 1), note number 64, pressure value 32.
func (pkpm *PolyphonicKeyPressureMessage) UnmarshalMIDI(b []byte) error {
	// check the number of bytes in the message
	if len(b) != PolyphonicKeyPressureMessageLength {
		return fmt.Errorf("polyphonic key pressure messages are made up of %d bytes, received %d byte(s): %w", PolyphonicKeyPressureMessageLength, len(b), ErrUnmarshallingMessage)
	}

	// make sure this is a status byte with the proper MSB
	if !ByteHasStatusMSB(b[0]) {
		return fmt.Errorf("polyphonic key pressure messages must have a status MSB: %w", ErrUnmarshallingMessage)
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

	// form the pressure
	pressure := NewPressureFromByte(b[2])

	*pkpm = PolyphonicKeyPressureMessage{
		Channel:  channel,
		Note:     note,
		Pressure: pressure,
	}
	return nil
}

// UnmarshalRunningStatusMIDI unmarshalls raw bytes into a PolyphonicKeyPressureMessage struct pointer. Note-On running status messages are
// represented by two bytes (left to right): note number, note velocity.
//
// Example: []byte{0b01000000, 0b00100000}
//
// The example forms a Note-On running status message for note number 64, velocity value 32.
func (pkpm *PolyphonicKeyPressureMessage) UnmarshalRunningStatusMIDI(b []byte) error {
	// check the number of bytes in the running status message
	if len(b) != PolyphonicKeyPressureMessageLength-1 {
		return fmt.Errorf("polyphonic key pressure running status messages are made up of %d bytes, received %d byte(s): %w", PolyphonicKeyPressureMessageLength-1, len(b), ErrUnmarshallingMessage)
	}

	// form the note number
	note, err := NewNoteFromByte(b[0])
	if err != nil {
		return fmt.Errorf("invalid note number %#v (%v) from running status note byte: %w", b[0], err, ErrUnmarshallingMessage)
	}

	// form the pressure
	pressure := NewPressureFromByte(b[1])

	*pkpm = PolyphonicKeyPressureMessage{
		Note:     note,
		Pressure: pressure,
	}
	return nil
}

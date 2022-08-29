package midiv1

import "fmt"

const (
	// ChannelPressureMessageStatusCode represents the message code within the status nibble
	ChannelPressureMessageCode Nibble = 0b01010000

	// ChannelPressureMessageLength represents the number of bytes in a full Channel Pressure message.
	ChannelPressureMessageLength int = 3

	// ChannelPressureMessageStatusNibble represents the status nibble within the status byte
	ChannelPressureMessageStatusNibble Status = Status(StatusMessageMSB) | Status(ChannelPressureMessageCode)
)

// ChannelPressureMessage represents a Channel Pressure Channel Voice message.
type ChannelPressureMessage struct {
	// Channel represents the channel number where this message will be sent.
	Channel Channel

	// Note represents the note number that will be sent with this message.
	Note Note

	// Pressure represents the relative applied pressure of the note in this message.
	Pressure Pressure
}

// GetMessageName returns the name of this Channel Pressure message.
func (cpm *ChannelPressureMessage) GetMessageName() string {
	return "Channel Pressure"
}

// MarshalMIDI marshalls a ChannelPressureMessage MIDI message into its raw bytes
func (cpm ChannelPressureMessage) MarshalMIDI() ([]byte, error) {
	return []byte{
		MakeStatusByte(ChannelPressureMessageStatusNibble, cpm.Channel),
		byte(cpm.Note),
		byte(cpm.Pressure),
	}, nil
}

// MarshalRunningStatusMIDI marshalls a running status MIDI message into its raw bytes.
func (cpm ChannelPressureMessage) MarshalRunningStatusMIDI() ([]byte, error) {
	return []byte{
		byte(cpm.Note),
		byte(cpm.Pressure),
	}, nil
}

// UnmarshalMIDI unmarshalls raw bytes into a ChannelPressureMessage struct pointer. Channel Pressure messages are
// represented by three bytes (left to right): status/channel, note number, note pressure.
//
// Example: []byte{0b11010001, 0b01000000, 0b00100000}
//
// The example forms a Channel Pressure message for channel 2 (index 1), note number 64, pressure value 32.
func (cpm *ChannelPressureMessage) UnmarshalMIDI(b []byte) error {
	// check the number of bytes in the message
	if len(b) != ChannelPressureMessageLength {
		return fmt.Errorf("channel pressure messages are made up of %d bytes, received %d byte(s): %w", ChannelPressureMessageLength, len(b), ErrUnmarshallingMessage)
	}

	// make sure this is a status byte with the proper MSB
	if !ByteHasStatusMSB(b[0]) {
		return fmt.Errorf("channel pressure messages must have a status MSB: %w", ErrUnmarshallingMessage)
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

	*cpm = ChannelPressureMessage{
		Channel:  channel,
		Note:     note,
		Pressure: pressure,
	}
	return nil
}

// UnmarshalRunningStatusMIDI unmarshalls raw bytes into a ChannelPressureMessage struct pointer. Channel Pressure running status messages are
// represented by two bytes (left to right): note number, note pressure.
//
// Example: []byte{0b01000000, 0b00100000}
//
// The example forms a Channel Pressure running status message for note number 64, pressure value 32.
func (cpm *ChannelPressureMessage) UnmarshalRunningStatusMIDI(b []byte) error {
	// check the number of bytes in the running status message
	if len(b) != ChannelPressureMessageLength-1 {
		return fmt.Errorf("channel pressure running status messages are made up of %d bytes, received %d byte(s): %w", ChannelPressureMessageLength-1, len(b), ErrUnmarshallingMessage)
	}

	// form the note number
	note, err := NewNoteFromByte(b[0])
	if err != nil {
		return fmt.Errorf("invalid note number %#v (%v) from running status note byte: %w", b[0], err, ErrUnmarshallingMessage)
	}

	// form the pressure
	pressure := NewPressureFromByte(b[1])

	*cpm = ChannelPressureMessage{
		Note:     note,
		Pressure: pressure,
	}
	return nil
}

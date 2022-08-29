package midiv1

import "fmt"

const (
	// PitchBendChangeMessageStatusCode represents the message code within the status nibble
	PitchBendChangeMessageCode Nibble = 0b01110000

	// PitchBendChangeMessageLength represents the number of bytes in a full Pitch Bend Change message.
	PitchBendChangeMessageLength int = 3

	// PitchBendChangeMessageStatusNibble represents the status nibble within the status byte
	PitchBendChangeMessageStatusNibble Status = Status(StatusMessageMSB) | Status(PitchBendChangeMessageCode)

	// PitchBendChangeMessageStringFormat represents the printf-compatible format specifically for a Pitch Bend Change message string.
	PitchBendChangeMessageStringFormat string = "%s:%s:%d:%d"
)

// PitchBendChangeMessage represents a Pitch Bend Change Channel Voice message.
type PitchBendChangeMessage struct {
	// Channel represents the channel number where this message will be sent.
	Channel Channel

	// PitchBend represents the pitch bend change for this message.
	PitchBend PitchBend
}

// GetMessageName returns the name of this Pitch Bend Change message.
func (pbm *PitchBendChangeMessage) GetMessageName() string {
	return "Pitch Bend Change"
}

// MarshalMIDI marshalls a PitchBendChangeMessage MIDI message into its raw bytes
func (pbm PitchBendChangeMessage) MarshalMIDI() ([]byte, error) {
	return []byte{
		MakeStatusByte(PitchBendChangeMessageStatusNibble, pbm.Channel),
		pbm.PitchBend.GetLSB(),
		pbm.PitchBend.GetMSB(),
	}, nil
}

// MarshalRunningStatusMIDI marshalls a running status MIDI message into its raw bytes.
func (pbm PitchBendChangeMessage) MarshalRunningStatusMIDI() ([]byte, error) {
	return []byte{
		pbm.PitchBend.GetLSB(),
		pbm.PitchBend.GetMSB(),
	}, nil
}

// String returns the human-readable representation of the MIDI message.
func (pbm *PitchBendChangeMessage) String() string {
	return fmt.Sprintf(PitchBendChangeMessageStringFormat, MessageVersion, pbm.GetMessageName(), pbm.Channel, pbm.PitchBend)
}

// UnmarshalMIDI unmarshalls raw bytes into a PitchBendChangeMessage struct pointer. Pitch Bend Change messages are
// represented by three bytes (left to right): status/channel, pitch bend LSB, pitch bend MSB.
//
// Example: []byte{0b11110001, 0b11100010, 0b00011101}
//
// The example forms a Pitch Bend Change message for channel 2 (index 1), pitch bend value 7650 (LSB: E2, MSB: 1D).
func (pbm *PitchBendChangeMessage) UnmarshalMIDI(b []byte) error {
	// check the number of bytes in the message
	if len(b) != PitchBendChangeMessageLength {
		return fmt.Errorf("pitch bend change messages are made up of %d bytes, received %d byte(s): %w", PitchBendChangeMessageLength, len(b), ErrUnmarshallingMessage)
	}

	// make sure this is a status byte with the proper MSB
	if !ByteHasStatusMSB(b[0]) {
		return fmt.Errorf("pitch bend change messages must have a status MSB: %w", ErrUnmarshallingMessage)
	}

	// retrieve the channel nibble of the status byte to form the Channel value
	channel, err := ParseChannelFromStatusByte(b[0])
	if err != nil {
		return err
	}

	// form the pitch bend
	pitchBend := NewPitchBendFromBytes(b[2], b[1])

	*pbm = PitchBendChangeMessage{
		Channel:   channel,
		PitchBend: pitchBend,
	}
	return nil
}

// UnmarshalRunningStatusMIDI unmarshalls raw bytes into a PitchBendChangeMessage struct pointer. Pitch Bend Change running status messages are
// represented by two bytes (left to right): pitch bend LSB, pitch bend MSB.
//
// Example: []byte{0b11100010, 0b00011101}
//
// The example forms a Pitch Bend Change running status message for pitch bend value 7650 (LSB: E2, MSB: 1D).
func (pbm *PitchBendChangeMessage) UnmarshalRunningStatusMIDI(b []byte) error {
	// check the number of bytes in the running status message
	if len(b) != PitchBendChangeMessageLength-1 {
		return fmt.Errorf("pitch bend change running status messages are made up of %d bytes, received %d byte(s): %w", PitchBendChangeMessageLength-1, len(b), ErrUnmarshallingMessage)
	}

	// form the pitch bend
	pitchBend := NewPitchBendFromBytes(b[1], b[0])

	*pbm = PitchBendChangeMessage{
		PitchBend: pitchBend,
	}
	return nil
}

package midiv1

import "fmt"

const (
	// PitchBendMessageStatusCode represents the message code within the status nibble
	PitchBendMessageCode Nibble = 0b01110000

	// PitchBendMessageLength represents the number of bytes in a full Pitch Bend message.
	PitchBendMessageLength int = 3

	// PitchBendMessageStatusNibble represents the status nibble within the status byte
	PitchBendMessageStatusNibble Status = Status(StatusMessageMSB) | Status(PitchBendMessageCode)

	// PitchBendMessageStringFormat represents the printf-compatible format specifically for a Pitch Bend message string.
	PitchBendMessageStringFormat string = "%s:%s:%d:%d"
)

// PitchBendMessage represents a Pitch Bend Channel Voice message.
type PitchBendMessage struct {
	// Channel represents the channel number where this message will be sent.
	Channel Channel

	// PitchBend represents the pitch bend change for this message.
	PitchBend PitchBend
}

// GetMessageName returns the name of this Pitch Bend message.
func (pbm *PitchBendMessage) GetMessageName() string {
	return "Pitch Bend"
}

// MarshalMIDI marshalls a PitchBendMessage MIDI message into its raw bytes
func (pbm PitchBendMessage) MarshalMIDI() ([]byte, error) {
	return []byte{
		MakeStatusByte(PitchBendMessageStatusNibble, pbm.Channel),
		pbm.PitchBend.GetLSB(),
		pbm.PitchBend.GetMSB(),
	}, nil
}

// MarshalRunningStatusMIDI marshalls a running status MIDI message into its raw bytes.
func (pbm PitchBendMessage) MarshalRunningStatusMIDI() ([]byte, error) {
	return []byte{
		pbm.PitchBend.GetLSB(),
		pbm.PitchBend.GetMSB(),
	}, nil
}

// String returns the human-readable representation of the MIDI message.
func (pbm *PitchBendMessage) String() string {
	return fmt.Sprintf(PitchBendMessageStringFormat, MessageVersion, pbm.GetMessageName(), pbm.Channel, pbm.PitchBend)
}

// UnmarshalMIDI unmarshalls raw bytes into a PitchBendMessage struct pointer. Pitch Bend messages are
// represented by three bytes (left to right): status/channel, pitch bend LSB, pitch bend MSB.
//
// Example: []byte{0b11110001, 0b11100010, 0b00011101}
//
// The example forms a Pitch Bend message for channel 2 (index 1), pitch bend value 7650 (LSB: E2, MSB: 1D).
func (pbm *PitchBendMessage) UnmarshalMIDI(b []byte) error {
	// check the number of bytes in the message
	if len(b) != PitchBendMessageLength {
		return fmt.Errorf("pitch bend messages are made up of %d bytes, received %d byte(s): %w", PitchBendMessageLength, len(b), ErrUnmarshallingMessage)
	}

	// make sure this is a status byte with the proper MSB
	if !ByteHasStatusMSB(b[0]) {
		return fmt.Errorf("pitch bend messages must have a status MSB: %w", ErrUnmarshallingMessage)
	}

	// retrieve the channel nibble of the status byte to form the Channel value
	channel, err := ParseChannelFromStatusByte(b[0])
	if err != nil {
		return err
	}

	// form the pitch bend
	pitchBend := NewPitchBendFromBytes(b[2], b[1])

	*pbm = PitchBendMessage{
		Channel:   channel,
		PitchBend: pitchBend,
	}
	return nil
}

// UnmarshalRunningStatusMIDI unmarshalls raw bytes into a PitchBendMessage struct pointer. Pitch Bend running status messages are
// represented by two bytes (left to right): pitch bend LSB, pitch bend MSB.
//
// Example: []byte{0b11100010, 0b00011101}
//
// The example forms a Pitch Bend running status message for pitch bend value 7650 (LSB: E2, MSB: 1D).
func (pbm *PitchBendMessage) UnmarshalRunningStatusMIDI(b []byte) error {
	// check the number of bytes in the running status message
	if len(b) != PitchBendMessageLength-1 {
		return fmt.Errorf("pitch bend running status messages are made up of %d bytes, received %d byte(s): %w", PitchBendMessageLength-1, len(b), ErrUnmarshallingMessage)
	}

	// form the pitch bend
	pitchBend := NewPitchBendFromBytes(b[1], b[0])

	*pbm = PitchBendMessage{
		PitchBend: pitchBend,
	}
	return nil
}

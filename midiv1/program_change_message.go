package midiv1

import "fmt"

const (
	// ProgramChangeMessageStatusCode represents the message code within the status nibble
	ProgramChangeMessageCode Nibble = 0b01000000

	// ProgramChangeMessageLength represents the number of bytes in a full Program Change message.
	ProgramChangeMessageLength int = 2

	// ProgramChangeMessageStatusNibble represents the status nibble within the status byte
	ProgramChangeMessageStatusNibble Status = Status(StatusMessageMSB) | Status(ProgramChangeMessageCode)

	// ProgramChangeMessageStringFormat represents the printf-compatible format specifically for a Program Change message string.
	ProgramChangeMessageStringFormat string = "%s:%s:%d:%d"
)

// ProgramChangeMessage represents a Program Change Channel Voice message.
type ProgramChangeMessage struct {
	// Channel represents the channel number where this message will be sent.
	Channel Channel

	// Program represents the program for this message.
	Program Program
}

// GetMessageName returns the name of this Program Change message.
func (pcm *ProgramChangeMessage) GetMessageName() string {
	return "Program Change"
}

// MarshalMIDI marshalls a ProgramChangeMessage MIDI message into its raw bytes
func (pcm ProgramChangeMessage) MarshalMIDI() ([]byte, error) {
	return []byte{
		MakeStatusByte(ProgramChangeMessageStatusNibble, pcm.Channel),
		byte(pcm.Program),
	}, nil
}

// MarshalRunningStatusMIDI marshalls a running status MIDI message into its raw bytes.
func (pcm ProgramChangeMessage) MarshalRunningStatusMIDI() ([]byte, error) {
	return []byte{
		byte(pcm.Program),
	}, nil
}

// String returns the human-readable representation of the MIDI message.
func (pcm *ProgramChangeMessage) String() string {
	return fmt.Sprintf(ProgramChangeMessageStringFormat, MessageVersion, pcm.GetMessageName(), pcm.Channel, pcm.Program)
}

// UnmarshalMIDI unmarshalls raw bytes into a ProgramChangeMessage struct pointer. Program Change messages are
// represented by two bytes (left to right): status/channel, program number.
//
// Example: []byte{0b11000001, 0b01100010}
//
// The example forms a Program Change message for channel 2 (index 1), program number 98.
func (pcm *ProgramChangeMessage) UnmarshalMIDI(b []byte) error {
	// check the number of bytes in the message
	if len(b) != ProgramChangeMessageLength {
		return fmt.Errorf("program change messages are made up of %d bytes, received %d byte(s): %w", ProgramChangeMessageLength, len(b), ErrUnmarshallingMessage)
	}

	// make sure this is a status byte with the proper MSB
	if !ByteHasStatusMSB(b[0]) {
		return fmt.Errorf("program change messages must have a status MSB: %w", ErrUnmarshallingMessage)
	}

	// retrieve the channel nibble of the status byte to form the Channel value
	channel, err := ParseChannelFromStatusByte(b[0])
	if err != nil {
		return err
	}

	// form the program
	program, err := NewProgramFromByte(b[1])
	if err != nil {
		return fmt.Errorf("invalid program number (%v) from program byte: %w", err, ErrUnmarshallingMessage)
	}

	*pcm = ProgramChangeMessage{
		Channel: channel,
		Program: program,
	}
	return nil
}

// UnmarshalRunningStatusMIDI unmarshalls raw bytes into a ProgramChangeMessage struct pointer. Program Change running status messages are
// represented by one byte: program change.
//
// Example: []byte{0b01100010}
//
// The example forms a Program Change running status message for program number 98.
func (pcm *ProgramChangeMessage) UnmarshalRunningStatusMIDI(b []byte) error {
	// check the number of bytes in the running status message
	if len(b) != ProgramChangeMessageLength-1 {
		return fmt.Errorf("program change running status messages are made up of %d bytes, received %d byte(s): %w", ProgramChangeMessageLength-1, len(b), ErrUnmarshallingMessage)
	}

	// form the program
	program, err := NewProgramFromByte(b[0])
	if err != nil {
		return fmt.Errorf("invalid program number (%v) from program byte: %w", err, ErrUnmarshallingMessage)
	}

	*pcm = ProgramChangeMessage{
		Program: program,
	}
	return nil
}

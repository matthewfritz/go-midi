package midiv1

import (
	"errors"
	"fmt"
)

const (
	// MinPitchBend represents the value of the lowest pitch bend.
	MinPitchBend PitchBend = -8192

	// MaxPitchBend represents the value of the highest pitch bend.
	MaxPitchBend PitchBend = 8192

	// ZeroPitchBend represents the value for no pitch bend.
	ZeroPitchBend PitchBend = 0
)

// PitchBend represents the pitch bend value of an individual MIDI note. Valid values are between -8192 and 8192 inclusive
// when converted to an integer.
//
// PitchBend is only used in conjunction with Pitch Bend Channel Voice messages.
type PitchBend int16

var (
	// ErrInvalidPitchBend represents an invalid MIDI pitch bend value.
	ErrInvalidPitchBend error = errors.New("invalid MIDI pitch bend")
)

// NewPitchBend returns a PitchBend instance from an integer value.
func NewPitchBend(pitchBend int) (PitchBend, error) {
	if pitchBend < int(MinPitchBend) || pitchBend > int(MaxPitchBend) {
		return ZeroPitchBend, fmt.Errorf("valid pitch bend is between %d and %d, inclusive: %w", MinPitchBend, MaxPitchBend, ErrInvalidChannel)
	}
	return PitchBend(pitchBend), nil
}

// GetLSB returns the least-significant byte of the pitch bend value.
func (pb PitchBend) GetLSB() byte {
	// https://stackoverflow.com/a/17260211
	return byte(pb & 0xFF)
}

// GetMSB returns the most-significant byte of the pitch bend value.
func (pb PitchBend) GetMSB() byte {
	// https://stackoverflow.com/a/17260211
	return byte(pb >> 8)
}

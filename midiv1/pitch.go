package midiv1

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

// NewPitchBend returns a PitchBend instance from an integer value.
func NewPitchBend(pitchBend int) PitchBend {
	if pitchBend < int(MinPitchBend) {
		return MinPitchBend
	}
	if pitchBend > int(MaxPitchBend) {
		return MaxPitchBend
	}
	return PitchBend(pitchBend)
}

// NewPitchBendFromBytes returns a PitchBend instance from a most-significant byte and a least-significant byte.
func NewPitchBendFromBytes(msb byte, lsb byte) PitchBend {
	return NewPitchBend(int((int16(msb) << 8) | int16(lsb)))
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

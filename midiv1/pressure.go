package midiv1

// Pressure represents the pressure applied to an individual MIDI note. Valid values are between 0 and 127 inclusive.
type Pressure int

const (
	// ZeroPressure represents the lowest possible pressure value for a MIDI note (0% strength).
	ZeroPressure Pressure = 0

	// LowPressure represents a pressure value of 25% strength.
	LowPressure Pressure = 31

	// MiddlePressure represents a pressure value of 50% strength.
	MiddlePressure Pressure = 63

	// HighPressure represents a pressure value of 75% strength.
	HighPressure Pressure = 95

	// FullPressure represents the highest possible pressure value for a MIDI note (100% strength).
	FullPressure Pressure = 127
)

// NewPressure returns a Pressure based on the integer argument, clamped within the overall minimum and maximum values.
func NewPressure(pressure int) Pressure {
	if pressure < int(ZeroPressure) {
		return ZeroPressure
	}
	if pressure > int(FullPressure) {
		return FullPressure
	}
	return Pressure(pressure)
}

// NewPressureFromByte returns a Pressure based on the byte argument, clamped within the overall minimum and maximum values.
func NewPressureFromByte(pressure byte) Pressure {
	if pressure < byte(ZeroPressure) {
		return ZeroPressure
	}
	if pressure > byte(FullPressure) {
		return FullPressure
	}
	return Pressure(pressure)
}

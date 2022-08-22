package midi

// Velocity represents the strength of an individual MIDI note. Valid values are between 0 and 127 inclusive.
type Velocity int

const (
	// ZeroVelocity represents the lowest possible velocity value for a MIDI note (0% strength).
	ZeroVelocity Velocity = 0

	// LowVelocity represents a velocity value of 25% strength.
	LowVelocity Velocity = 31

	// MiddleVelocity represents a velocity value of 50% strength.
	MiddleVelocity Velocity = 63

	// HighVelocity represents a velocity value of 75% strength.
	HighVelocity Velocity = 95

	// FullVelocity represents the highest possible velocity value for a MIDI note (100% strength).
	FullVelocity Velocity = 127
)

// RandomVelocityInRange returns a random note velocity between the provided minimum and maximum values.
func RandomVelocityInRange(min Velocity, max Velocity) Velocity {
	return 0
}

// RandomVelocity returns a random note velocity between the overall lowest and highest values.
func RandomVelocity() Velocity {
	return RandomVelocityInRange(ZeroVelocity, FullVelocity)
}

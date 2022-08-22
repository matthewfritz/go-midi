package midi

import "testing"

func Test_SafeRandomVelocityInRange(t *testing.T) {
	t.Parallel()

	velocityRandomizer := NewVelocityRandomizer()

	minVelocity := NewVelocity(61)
	maxVelocity := NewVelocity(63)

	got := velocityRandomizer.SafeRandomVelocityInRange(minVelocity, maxVelocity)
	if got < minVelocity || got > maxVelocity {
		t.Fatalf("expected random custom ranged velocity between %v and %v, got %v", minVelocity, maxVelocity, got)
	}
}

func Test_SafeRandomVelocity(t *testing.T) {
	t.Parallel()

	velocityRandomizer := NewVelocityRandomizer()

	got := velocityRandomizer.SafeRandomVelocity()
	if got < ZeroVelocity || got > FullVelocity {
		t.Fatalf("expected random ranged velocity between %v and %v, got %v", ZeroVelocity, FullVelocity, got)
	}
}

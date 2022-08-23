package midi

import (
	"errors"
	"testing"
)

func Test_SafeRandomVelocityInRange(t *testing.T) {
	t.Parallel()
	errorTests := map[string]struct {
		minVelocity           Velocity
		maxVelocity           Velocity
		expectedSentinelError error
	}{
		"minimum velocity greater than maximum velocity": {
			minVelocity:           NewVelocity(59),
			maxVelocity:           NewVelocity(53),
			expectedSentinelError: ErrRandomVelocity,
		},
	}

	velocityRandomizer := NewVelocityRandomizer()

	// sentinel error tests
	for name, test := range errorTests {
		t.Run(name, func(t *testing.T) {
			_, err := velocityRandomizer.SafeRandomVelocityInRange(test.minVelocity, test.maxVelocity)
			if err == nil {
				t.Fatalf("expected non-nil error (%v), got nil error", test.expectedSentinelError)
			}
			if !errors.Is(err, test.expectedSentinelError) {
				t.Fatalf("expected %v error, got %v", test.expectedSentinelError, err)
			}
		})
	}

	// test for no randomization with only one possible value in range
	sameVelocity := NewVelocity(12)
	got, err := velocityRandomizer.SafeRandomVelocityInRange(sameVelocity, sameVelocity)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got != sameVelocity {
		t.Fatalf("expected %v Velocity, got %v", sameVelocity, got)
	}

	// standard randomization within range
	minVelocity := NewVelocity(61)
	maxVelocity := NewVelocity(63)
	got, err = velocityRandomizer.SafeRandomVelocityInRange(minVelocity, maxVelocity)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got < minVelocity || got > maxVelocity {
		t.Fatalf("expected random custom ranged velocity between %v and %v, got %v", minVelocity, maxVelocity, got)
	}
}

func Test_SafeRandomVelocity(t *testing.T) {
	t.Parallel()

	velocityRandomizer := NewVelocityRandomizer()

	got, err := velocityRandomizer.SafeRandomVelocity()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got < ZeroVelocity || got > FullVelocity {
		t.Fatalf("expected random ranged velocity between %v and %v, got %v", ZeroVelocity, FullVelocity, got)
	}
}

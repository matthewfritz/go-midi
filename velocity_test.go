package midi

import "testing"

func Test_NewVelocity(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		velocityInt      int
		expectedVelocity Velocity
	}{
		"velocity clamps to zero value": {
			velocityInt:      -1,
			expectedVelocity: ZeroVelocity,
		},
		"velocity clamps to maximum value": {
			velocityInt:      128,
			expectedVelocity: FullVelocity,
		},
		"velocity is intended integer": {
			velocityInt:      53,
			expectedVelocity: 53,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewVelocity(test.velocityInt)
			if got != test.expectedVelocity {
				t.Fatalf("expected %v, got %v", test.expectedVelocity, got)
			}
		})
	}
}

func Test_RandomVelocityInRange(t *testing.T) {
	t.Parallel()

	// seed the source
	SeedRandomVelocitySource()

	minVelocity := NewVelocity(61)
	maxVelocity := NewVelocity(63)

	got := RandomVelocityInRange(minVelocity, maxVelocity)
	if got < minVelocity || got > maxVelocity {
		t.Fatalf("expected random custom ranged velocity between %v and %v, got %v", minVelocity, maxVelocity, got)
	}
}

func Test_RandomVelocity(t *testing.T) {
	t.Parallel()

	// seed the source
	SeedRandomVelocitySource()

	got := RandomVelocity()
	if got < ZeroVelocity || got > FullVelocity {
		t.Fatalf("expected random ranged velocity between %v and %v, got %v", ZeroVelocity, FullVelocity, got)
	}
}

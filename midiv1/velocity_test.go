package midiv1

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

func Test_NewVelocityFromByte(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		velocityByte     byte
		expectedVelocity Velocity
	}{
		"velocity clamps to maximum value": {
			velocityByte:     128,
			expectedVelocity: FullVelocity,
		},
		"velocity is intended byte": {
			velocityByte:     53,
			expectedVelocity: 53,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewVelocityFromByte(test.velocityByte)
			if got != test.expectedVelocity {
				t.Fatalf("expected %v, got %v", test.expectedVelocity, got)
			}
		})
	}
}

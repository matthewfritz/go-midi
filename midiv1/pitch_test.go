package midiv1

import "testing"

func Test_NewPitchBend(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		pitchBendInt      int
		expectedPitchBend PitchBend
	}{
		"pitch bend clamps to minimum value": {
			pitchBendInt:      -8913,
			expectedPitchBend: MinPitchBend,
		},
		"pitch bend clamps to maximum value": {
			pitchBendInt:      8913,
			expectedPitchBend: MaxPitchBend,
		},
		"pitch bend is intended integer": {
			pitchBendInt:      53,
			expectedPitchBend: 53,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewPitchBend(test.pitchBendInt)
			if got != test.expectedPitchBend {
				t.Fatalf("expected %v, got %v", test.expectedPitchBend, got)
			}
		})
	}
}

func Test_NewPitchBendFromBytes(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		pitchBendBytes    [2]byte
		expectedPitchBend PitchBend
	}{
		"pitch bend MSB and LSB make up -8192": {
			pitchBendBytes:    [2]byte{0xE0, 0x00},
			expectedPitchBend: MinPitchBend,
		},
		"pitch bend MSB and LSB make up 8192": {
			pitchBendBytes:    [2]byte{0x20, 0x00},
			expectedPitchBend: MaxPitchBend,
		},
		"pitch bend MSB and LSB make up 53": {
			pitchBendBytes:    [2]byte{0x00, 0x35},
			expectedPitchBend: 53,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewPitchBendFromBytes(test.pitchBendBytes[0], test.pitchBendBytes[1])
			if got != test.expectedPitchBend {
				t.Fatalf("expected %v, got %v", test.expectedPitchBend, got)
			}
		})
	}
}

func Test_PitchBend_GetLSB(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		pitchBend   PitchBend
		expectedLSB byte
	}{
		"pitch bend LSB is 0x00": {
			pitchBend:   MinPitchBend,
			expectedLSB: 0x00,
		},
		"pitch bend LSB is 0x35": {
			pitchBend:   53,
			expectedLSB: 0x35,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.pitchBend.GetLSB()
			if got != test.expectedLSB {
				t.Fatalf("expected %v, got %v", test.expectedLSB, got)
			}
		})
	}
}

func Test_PitchBend_GetMSB(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		pitchBend   PitchBend
		expectedMSB byte
	}{
		"pitch bend MSB is 0xE0": {
			pitchBend:   MinPitchBend,
			expectedMSB: 0xE0,
		},
		"pitch bend MSB is 0x20": {
			pitchBend:   MaxPitchBend,
			expectedMSB: 0x20,
		},
		"pitch bend MSB is 0x00": {
			pitchBend:   53,
			expectedMSB: 0x00,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.pitchBend.GetMSB()
			if got != test.expectedMSB {
				t.Fatalf("expected %v, got %v", test.expectedMSB, got)
			}
		})
	}
}

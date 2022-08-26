package midiv1

import "testing"

func Test_NewPressure(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		pressureInt      int
		expectedPressure Pressure
	}{
		"pressure clamps to zero value": {
			pressureInt:      -1,
			expectedPressure: ZeroPressure,
		},
		"pressure clamps to maximum value": {
			pressureInt:      128,
			expectedPressure: FullPressure,
		},
		"pressure is intended integer": {
			pressureInt:      53,
			expectedPressure: 53,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewPressure(test.pressureInt)
			if got != test.expectedPressure {
				t.Fatalf("expected %v, got %v", test.expectedPressure, got)
			}
		})
	}
}

func Test_NewPressureFromByte(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		pressureByte     byte
		expectedPressure Pressure
	}{
		"pressure clamps to maximum value": {
			pressureByte:     128,
			expectedPressure: FullPressure,
		},
		"pressure is intended byte": {
			pressureByte:     53,
			expectedPressure: 53,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewPressureFromByte(test.pressureByte)
			if got != test.expectedPressure {
				t.Fatalf("expected %v, got %v", test.expectedPressure, got)
			}
		})
	}
}

package midiv1

import (
	"errors"
	"testing"
)

func Test_NewProgram(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		programInt      int
		expectedProgram Program
		err             error
	}{
		"program number too low": {
			programInt:      -1,
			expectedProgram: MinProgram,
			err:             ErrInvalidProgram,
		},
		"program number too high": {
			programInt:      128,
			expectedProgram: MinProgram,
			err:             ErrInvalidProgram,
		},
		"program is intended value": {
			programInt:      12,
			expectedProgram: 12,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewProgram(test.programInt)
			if test.err == nil && err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
			if test.err != nil {
				if err == nil {
					t.Fatalf("expected non-nil %v error, got nil error", test.err)
				}
				if !errors.Is(err, test.err) {
					t.Fatalf("expected %v error, got %v", test.err, err)
				}
			}
			if got != test.expectedProgram {
				t.Fatalf("expected %v, got %v", test.expectedProgram, got)
			}
		})
	}
}

func Test_NewProgramFromByte(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		programByte     byte
		expectedProgram Program
		err             error
	}{
		"program number too high": {
			programByte:     128,
			expectedProgram: MinProgram,
			err:             ErrInvalidProgram,
		},
		"program is intended value": {
			programByte:     12,
			expectedProgram: 12,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewProgramFromByte(test.programByte)
			if test.err == nil && err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
			if test.err != nil {
				if err == nil {
					t.Fatalf("expected non-nil %v error, got nil error", test.err)
				}
				if !errors.Is(err, test.err) {
					t.Fatalf("expected %v error, got %v", test.err, err)
				}
			}
			if got != test.expectedProgram {
				t.Fatalf("expected %v, got %v", test.expectedProgram, got)
			}
		})
	}
}

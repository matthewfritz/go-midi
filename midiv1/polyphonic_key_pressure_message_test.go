package midiv1

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func Test_PolyphonicKeyPressureMessage_MarshalMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		message  PolyphonicKeyPressureMessage
		expected []byte
	}{
		"message marshalls into expected bytes": {
			message: PolyphonicKeyPressureMessage{
				Channel:  1,
				Note:     64,
				Pressure: 32,
			},
			expected: []byte{0b10100001, 0b01000000, 0b00100000},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := test.message.MarshalMIDI()
			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
			if !bytes.Equal(test.expected, got) {
				t.Fatalf("expected %#v, got %#v", test.expected, got)
			}
		})
	}
}

func Test_PolyphonicKeyPressureMessage_MarshalRunningStatusMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		message  PolyphonicKeyPressureMessage
		expected []byte
	}{
		"running status message marshalls into expected bytes": {
			message: PolyphonicKeyPressureMessage{
				Note:     64,
				Pressure: 32,
			},
			expected: []byte{0b01000000, 0b00100000},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := test.message.MarshalRunningStatusMIDI()
			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
			if !bytes.Equal(test.expected, got) {
				t.Fatalf("expected %#v, got %#v", test.expected, got)
			}
		})
	}
}

func Test_PolyphonicKeyPressureMessage_UnmarshalMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		b               []byte
		expectedMessage PolyphonicKeyPressureMessage
		err             error
	}{
		"byte slice is not proper length": {
			b:   []byte{0b10100001, 0b01000000},
			err: ErrUnmarshallingMessage,
		},
		"first byte does not have a status MSB": {
			b:   []byte{0b00100001, 0b01000000, 0b00100000},
			err: ErrUnmarshallingMessage,
		},
		"second byte is an invalid note number": {
			b:   []byte{0b10100001, 0b11000000, 0b00100000},
			err: ErrUnmarshallingMessage,
		},
		"bytes unmarshal into expected message": {
			b: []byte{0b10100001, 0b01000000, 0b00100000},
			expectedMessage: PolyphonicKeyPressureMessage{
				Channel:  1,
				Note:     64,
				Pressure: 32,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got PolyphonicKeyPressureMessage
			err := (&got).UnmarshalMIDI(test.b)
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
			if !reflect.DeepEqual(test.expectedMessage, got) {
				t.Fatalf("expected %+v, got %+v", test.expectedMessage, got)
			}
		})
	}
}

func Test_PolyphonicKeyPressureMessage_UnmarshalRunningStatusMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		b               []byte
		expectedMessage PolyphonicKeyPressureMessage
		err             error
	}{
		"byte slice is not proper length": {
			b:   []byte{0b10100001, 0b10010001, 0b01000000},
			err: ErrUnmarshallingMessage,
		},
		"second byte is an invalid note number": {
			b:   []byte{0b11000000, 0b00100000},
			err: ErrUnmarshallingMessage,
		},
		"bytes unmarshal into expected message": {
			b: []byte{0b01000000, 0b00100000},
			expectedMessage: PolyphonicKeyPressureMessage{
				Note:     64,
				Pressure: 32,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got PolyphonicKeyPressureMessage
			err := (&got).UnmarshalRunningStatusMIDI(test.b)
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
			if !reflect.DeepEqual(test.expectedMessage, got) {
				t.Fatalf("expected %+v, got %+v", test.expectedMessage, got)
			}
		})
	}
}

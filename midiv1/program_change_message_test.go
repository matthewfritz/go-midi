package midiv1

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_ProgramChangeMessage_GetMessageName(t *testing.T) {
	t.Parallel()
	message := ProgramChangeMessage{}
	expected := "Program Change"
	if message.GetMessageName() != expected {
		t.Fatalf("expected %s, got %s", expected, message.GetMessageName())
	}
}

func Test_ProgramChangeMessage_MarshalMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		message  ProgramChangeMessage
		expected []byte
	}{
		"message marshalls into expected bytes": {
			message: ProgramChangeMessage{
				Channel: 1,
				Program: 64,
			},
			expected: []byte{0b11000001, 0b01000000},
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

func Test_ProgramChangeMessage_MarshalRunningStatusMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		message  ProgramChangeMessage
		expected []byte
	}{
		"running status message marshalls into expected bytes": {
			message: ProgramChangeMessage{
				Program: 64,
			},
			expected: []byte{0b01000000},
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

func Test_ProgramChangeMessage_String(t *testing.T) {
	t.Parallel()
	message := ProgramChangeMessage{
		Channel: 1,
		Program: 64,
	}
	expected := fmt.Sprintf("%s:%s:%d:%d", MessageVersion, "Program Change", 1, 64)
	if message.String() != expected {
		t.Fatalf("expected %s, got %s", expected, message.String())
	}
}

func Test_ProgramChangeMessage_UnmarshalMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		b               []byte
		expectedMessage ProgramChangeMessage
		err             error
	}{
		"byte slice is not proper length": {
			b:   []byte{0b11000001, 0b01000000, 0b01000000},
			err: ErrUnmarshallingMessage,
		},
		"first byte does not have a status MSB": {
			b:   []byte{0b01000001, 0b01000000},
			err: ErrUnmarshallingMessage,
		},
		"second byte is an invalid program number": {
			b:   []byte{0b11000001, 0b10000001},
			err: ErrUnmarshallingMessage,
		},
		"bytes unmarshal into expected message": {
			b: []byte{0b11000001, 0b01000000},
			expectedMessage: ProgramChangeMessage{
				Channel: 1,
				Program: 64,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got ProgramChangeMessage
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

func Test_ProgramChangeMessage_UnmarshalRunningStatusMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		b               []byte
		expectedMessage ProgramChangeMessage
		err             error
	}{
		"byte slice is not proper length": {
			b:   []byte{0b11000001, 0b01000000},
			err: ErrUnmarshallingMessage,
		},
		"first byte is an invalid program number": {
			b:   []byte{0b10000001},
			err: ErrUnmarshallingMessage,
		},
		"bytes unmarshal into expected message": {
			b: []byte{0b01000000},
			expectedMessage: ProgramChangeMessage{
				Program: 64,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got ProgramChangeMessage
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

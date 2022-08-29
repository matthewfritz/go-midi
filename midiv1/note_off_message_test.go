package midiv1

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_NotOffMessage_GetMessageName(t *testing.T) {
	t.Parallel()
	message := NoteOffMessage{}
	expected := "Note-Off"
	if message.GetMessageName() != expected {
		t.Fatalf("expected %s, got %s", expected, message.GetMessageName())
	}
}

func Test_NoteOffMessage_MarshalMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		message  NoteOffMessage
		expected []byte
	}{
		"message marshalls into expected bytes": {
			message: NoteOffMessage{
				Channel:  1,
				Note:     64,
				Velocity: 32,
			},
			expected: []byte{0b10000001, 0b01000000, 0b00100000},
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

func Test_NoteOffMessage_MarshalRunningStatusMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		message  NoteOffMessage
		expected []byte
	}{
		"running status message marshalls into expected bytes": {
			message: NoteOffMessage{
				Note:     64,
				Velocity: 32,
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

func Test_NoteOffMessage_String(t *testing.T) {
	t.Parallel()
	message := NoteOffMessage{
		Channel:  1,
		Note:     64,
		Velocity: 32,
	}
	expected := fmt.Sprintf("%s:%s:%d:%d:%d", MessageVersion, "Note-Off", 1, 64, 32)
	if message.String() != expected {
		t.Fatalf("expected %s, got %s", expected, message.String())
	}
}

func Test_NoteOffMessage_UnmarshalMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		b               []byte
		expectedMessage NoteOffMessage
		err             error
	}{
		"byte slice is not proper length": {
			b:   []byte{0b10000001, 0b01000000},
			err: ErrUnmarshallingMessage,
		},
		"first byte does not have a status MSB": {
			b:   []byte{0b00000001, 0b01000000, 0b00100000},
			err: ErrUnmarshallingMessage,
		},
		"second byte is an invalid note number": {
			b:   []byte{0b10000001, 0b11000000, 0b00100000},
			err: ErrUnmarshallingMessage,
		},
		"bytes unmarshal into expected message": {
			b: []byte{0b10000001, 0b01000000, 0b00100000},
			expectedMessage: NoteOffMessage{
				Channel:  1,
				Note:     64,
				Velocity: 32,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got NoteOffMessage
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

func Test_NoteOffMessage_UnmarshalRunningStatusMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		b               []byte
		expectedMessage NoteOffMessage
		err             error
	}{
		"byte slice is not proper length": {
			b:   []byte{0b10000001, 0b10010001, 0b01000000},
			err: ErrUnmarshallingMessage,
		},
		"second byte is an invalid note number": {
			b:   []byte{0b11000000, 0b00100000},
			err: ErrUnmarshallingMessage,
		},
		"bytes unmarshal into expected message": {
			b: []byte{0b01000000, 0b00100000},
			expectedMessage: NoteOffMessage{
				Note:     64,
				Velocity: 32,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got NoteOffMessage
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

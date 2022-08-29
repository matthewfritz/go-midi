package midiv1

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func Test_PitchBendMessage_GetMessageName(t *testing.T) {
	message := PitchBendMessage{}
	expected := "Pitch Bend"
	if message.GetMessageName() != expected {
		t.Fatalf("expected %s, got %s", expected, message.GetMessageName())
	}
}

func Test_PitchBendMessage_MarshalMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		message  PitchBendMessage
		expected []byte
	}{
		"message marshalls into expected bytes": {
			message: PitchBendMessage{
				Channel:   1,
				PitchBend: 7650,
			},
			expected: []byte{0b11110001, 0b11100010, 0b00011101},
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

func Test_PitchBendMessage_MarshalRunningStatusMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		message  PitchBendMessage
		expected []byte
	}{
		"running status message marshalls into expected bytes": {
			message: PitchBendMessage{
				PitchBend: 7650,
			},
			expected: []byte{0b11100010, 0b00011101},
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

func Test_PitchBendMessage_UnmarshalMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		b               []byte
		expectedMessage PitchBendMessage
		err             error
	}{
		"byte slice is not proper length": {
			b:   []byte{0b11110001, 0b01000000},
			err: ErrUnmarshallingMessage,
		},
		"first byte does not have a status MSB": {
			b:   []byte{0b01110001, 0b01000000, 0b00100000},
			err: ErrUnmarshallingMessage,
		},
		"bytes unmarshal into expected message": {
			b: []byte{0b11010001, 0b11100010, 0b00011101},
			expectedMessage: PitchBendMessage{
				Channel:   1,
				PitchBend: 7650,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got PitchBendMessage
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

func Test_PitchBendMessage_UnmarshalRunningStatusMIDI(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		b               []byte
		expectedMessage PitchBendMessage
		err             error
	}{
		"byte slice is not proper length": {
			b:   []byte{0b11110001, 0b11100010, 0b00011101},
			err: ErrUnmarshallingMessage,
		},
		"bytes unmarshal into expected message": {
			b: []byte{0b11100010, 0b00011101},
			expectedMessage: PitchBendMessage{
				PitchBend: 7650,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got PitchBendMessage
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

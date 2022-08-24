package midiv1

import (
	"errors"
	"testing"
)

func Test_NewChannel(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		channelInt      int
		expectedChannel Channel
		err             error
	}{
		"channel number too low": {
			channelInt:      -1,
			expectedChannel: MinChannel,
			err:             ErrInvalidChannel,
		},
		"channel number too high": {
			channelInt:      17,
			expectedChannel: MinChannel,
			err:             ErrInvalidChannel,
		},
		"channel is intended value": {
			channelInt:      12,
			expectedChannel: 12,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewChannel(test.channelInt)
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
			if got != test.expectedChannel {
				t.Fatalf("expected %v, got %v", test.expectedChannel, got)
			}
		})
	}
}

func Test_NewChannelFromByte(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		channelByte     byte
		expectedChannel Channel
		err             error
	}{
		"channel number too high": {
			channelByte:     17,
			expectedChannel: MinChannel,
			err:             ErrInvalidChannel,
		},
		"channel is intended value": {
			channelByte:     12,
			expectedChannel: 12,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewChannelFromByte(test.channelByte)
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
			if got != test.expectedChannel {
				t.Fatalf("expected %v, got %v", test.expectedChannel, got)
			}
		})
	}
}

func Test_NewNote(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		noteInt      int
		expectedNote Note
		err          error
	}{
		"note number too low": {
			noteInt:      -1,
			expectedNote: MinNote,
			err:          ErrInvalidNote,
		},
		"note number too high": {
			noteInt:      128,
			expectedNote: MinNote,
			err:          ErrInvalidNote,
		},
		"note is intended value": {
			noteInt:      12,
			expectedNote: 12,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewNote(test.noteInt)
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
			if got != test.expectedNote {
				t.Fatalf("expected %v, got %v", test.expectedNote, got)
			}
		})
	}
}

func Test_NewNoteFromByte(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		noteByte     byte
		expectedNote Note
		err          error
	}{
		"note number too high": {
			noteByte:     128,
			expectedNote: MinNote,
			err:          ErrInvalidNote,
		},
		"note is intended value": {
			noteByte:     12,
			expectedNote: 12,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewNoteFromByte(test.noteByte)
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
			if got != test.expectedNote {
				t.Fatalf("expected %v, got %v", test.expectedNote, got)
			}
		})
	}
}

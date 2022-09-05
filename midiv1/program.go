package midiv1

import (
	"errors"
	"fmt"
)

// Program represents the program number of an individual MIDI message. Valid values are between 0 and 127 inclusive
// when converted to an integer.
//
// Program is only used in conjunction with Program Change Channel Voice messages.
type Program int

var (
	// ErrInvalidProgram represents an invalid MIDI program number.
	ErrInvalidProgram error = errors.New("invalid MIDI program number")
)

const (
	// MinProgram is the lowest MIDI program number available.
	MinProgram Program = 0

	// MaxProgram is the highest MIDI program number available.
	MaxProgram Program = 127
)

// NewProgram returns a Program based on the integer argument.
func NewProgram(program int) (Program, error) {
	if program < int(MinProgram) || program > int(MaxProgram) {
		return MinProgram, fmt.Errorf("valid program numbers are between %d and %d, inclusive: %w", MinProgram, MaxProgram, ErrInvalidProgram)
	}
	return Program(program), nil
}

// NewProgramFromByte returns a Program based on the byte argument.
func NewProgramFromByte(program byte) (Program, error) {
	if program < byte(MinProgram) || program > byte(MaxProgram) {
		return MinProgram, fmt.Errorf("valid program numbers are between %d and %d, inclusive: %w", MinProgram, MaxProgram, ErrInvalidProgram)
	}
	return Program(program), nil
}

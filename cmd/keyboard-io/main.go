package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/matthewfritz/go-midi/midiv1"
)

func main() {
	velocityRandomizer := midiv1.NewVelocityRandomizer()

	// expect a note number and velocity flag
	noteNumPtr := flag.Int("note", 0, "MIDI note number")
	velocityNumPtr := flag.Int("vel", 0, "MIDI note velocity value")
	flag.Parse()

	fmt.Printf("MIDI note number: %d\n", *noteNumPtr)
	fmt.Printf("MIDI note velocity: %d\n", midiv1.NewVelocity(*velocityNumPtr))

	randomVel, err := velocityRandomizer.SafeRandomVelocity()
	if err != nil {
		fmt.Printf("Error generating safe random velocity: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Random MIDI note velocity: %d\n", randomVel)
}

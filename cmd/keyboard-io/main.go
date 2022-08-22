package main

import (
	"fmt"

	"github.com/matthewfritz/go-midi"
)

func main() {
	midi.SeedRandomVelocitySource()
	fmt.Println(midi.RandomVelocity())
}

package main

import (
	"fmt"

	"github.com/matthewfritz/go-midi"
)

func main() {
	velocityRandomizer := midi.NewVelocityRandomizer()
	fmt.Println(velocityRandomizer.SafeRandomVelocity())
}

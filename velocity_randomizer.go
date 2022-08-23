package midi

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	// ErrRandomVelocity is returned when an error occurs generating a random velocity value
	ErrRandomVelocity error = errors.New("invalid random velocity")
)

// VelocityRandomizer generates random velocity values for MIDI notes.
type VelocityRandomizer struct {
	// mutex adds locking and unlocking around retrieving random integers from the randomizer
	mutex sync.RWMutex

	// randomizer is not concurrency-safe when using a non-default Source
	randomizer *rand.Rand
}

// NewVelocityRandomizer returns a new VelocityRandomizer, seeded with a specific Source.
func NewVelocityRandomizer() VelocityRandomizer {
	return VelocityRandomizer{
		randomizer: rand.New(rand.NewSource(time.Now().UnixMicro())),
	}
}

// RandomVelocityInRange returns a random note velocity between the provided minimum and maximum values.
// This method is NOT concurrency-safe.
func (vr *VelocityRandomizer) RandomVelocityInRange(min Velocity, max Velocity) (Velocity, error) {
	minInt := int(min)
	maxInt := int(max)
	if minInt > maxInt {
		// prevent a panic
		return NewVelocity(0), fmt.Errorf("minimum velocity cannot be greater than maximum velocity: %w", ErrRandomVelocity)
	}
	if minInt == maxInt {
		// range only allows for a single possible value
		return NewVelocity(minInt), nil
	}
	randVelInt := vr.randomizer.Intn(maxInt-minInt) + minInt
	return NewVelocity(randVelInt), nil
}

// RandomVelocity returns a random note velocity between the overall lowest and highest values.
// This method is NOT concurrency-safe.
func (vr *VelocityRandomizer) RandomVelocity() (Velocity, error) {
	return vr.RandomVelocityInRange(ZeroVelocity, FullVelocity)
}

// SafeRandomVelocityInRange returns a random note velocity between the provided minimum and maximum values.
// This method is concurrency-safe.
func (vr *VelocityRandomizer) SafeRandomVelocityInRange(min Velocity, max Velocity) (Velocity, error) {
	vr.mutex.Lock()
	defer vr.mutex.Unlock()
	return vr.RandomVelocityInRange(min, max)
}

// SafeRandomVelocity returns a random note velocity between the overall lowest and highest values.
// This method is concurrency-safe.
func (vr *VelocityRandomizer) SafeRandomVelocity() (Velocity, error) {
	vr.mutex.Lock()
	defer vr.mutex.Unlock()
	return vr.RandomVelocity()
}

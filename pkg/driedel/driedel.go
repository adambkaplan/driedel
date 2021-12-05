package driedel

import (
	"fmt"
	"math/rand"
	"time"
)

type SpinResult int

const (
	Nun SpinResult = iota
	Gimel
	Hey
	Shin
)

func (s SpinResult) String() string {
	switch s {
	case Nun:
		return "נ Nun"
	case Gimel:
		return "ג Gimel"
	case Hey:
		return "ה Hey"
	case Shin:
		return "ש Shin"
	}
	return fmt.Sprintf("%d is not a spin!", s)
}

// Driedel represents the top used in the Hanukah driedel game.
type Driedel struct {
	randomizer *rand.Rand
}

// NewDriedel creates a new Driedel with the current UNIX time in nanoseconds as the initial seed.
func NewDriedel() *Driedel {
	return NewDriedelWithSeed(time.Now().UnixNano())
}

// NewDriedelWithSeed creates a new Driedel with the provided seed for the random number generator.
func NewDriedelWithSeed(seed int64) *Driedel {
	return &Driedel{
		randomizer: rand.New(rand.NewSource(seed)),
	}
}

// Spin simulates the spinning of the driedel, returning a result of נ Nun, ג Gimel, ה Hey, or ש Shin.
func (d *Driedel) Spin() SpinResult {
	return SpinResult(d.spinDriedel())
}

func (d *Driedel) spinDriedel() int {
	return d.randomizer.Intn(4)
}

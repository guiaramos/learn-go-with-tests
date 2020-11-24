package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// Sleeper  type
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper object
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// DefaultSleeper obj
type DefaultSleeper struct{}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

// Sleep creates wait time
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown counts until go
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

// Sleep with configurable time
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

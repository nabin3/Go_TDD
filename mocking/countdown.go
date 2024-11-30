package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord  = "Go!"
	startCount = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration time.Duration
	Csleep   func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.Csleep(c.Duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := startCount; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

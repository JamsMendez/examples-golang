package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

// test ...
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Production ..
type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
  Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
  s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write (p []byte) (n int , err error) {
  s.Calls = append(s.Calls, write)
  return
}

type ConfigurablesSleeper struct {
  duration time.Duration
  sleep func(time.Duration)
}

func (c *ConfigurablesSleeper) Sleep() {
  c.sleep(c.duration)
}


type SpyTime struct {
  durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
  s.durationSlept = duration
}

// [02]
/* func Countdown(writer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprint(writer, finalWord)
} */

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
		sleeper.Sleep()
	}

	fmt.Fprint(writer, finalWord)
}

func CountdownBreak(writer io.Writer, sleeper Sleeper) {
  for i := countdownStart; i > 0; i-- {
    sleeper.Sleep()
  }

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
	}

	fmt.Fprint(writer, finalWord)
}

package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// [01]
/* func TestCountDown(t *testing.T) {
  buffer := &bytes.Buffer{}

  Countdown(buffer)

  got := buffer.String()
  want := "3"

  if got != want {
    t.Errorf("expected %q want %q", got, want)
  }
} */

// [02]
/* func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("expected %q want %q", got, want)
	}
} */

func TestCountdonw(t *testing.T) {

	t.Run("prints 3 Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("expected %q want %q", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enought calls to slepper, want 3 got %d", spySleeper.Calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

	t.Run("test configurable sleeper", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := ConfigurablesSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}

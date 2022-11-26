package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
  buffer := bytes.Buffer{}
  Greet(&buffer, "JamsMendez")

  got := buffer.String()
  want := "Hello, JamsMendez"

  if got != want {
    t.Errorf("expected %q want %q", got, want)
  }
}

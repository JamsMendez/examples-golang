package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
  repeated := Repeat("a")
  expected := "aaaaa"
  
  if repeated != expected {
    t.Errorf("expected %q but got %q", expected, repeated)
  }
}

func ExampleRepeat() {
  repeated := Repeat("a")
  fmt.Println(repeated)
  // Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a")
  }
}

// go test -bench="."

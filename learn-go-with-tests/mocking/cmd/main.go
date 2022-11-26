package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
  Countdown(os.Stdout)
}

func Countdown(writer io.Writer) {
  fmt.Fprint(writer, 3)
}

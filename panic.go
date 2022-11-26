package main

import "fmt"

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Call panic!")
		// The control flow is interrupted, and all the deferred functions start to print
		// the Deferred message.

		panic("highlow() low greater that high")
	}

	defer fmt.Printf("Deferred highlow(%d, %d)\n", high, low)

	fmt.Printf("Call: highlow(%d, %d)\n", high, low)

	highlow(high, low+1)
}

func panicCall() {
	highlow(2, 0)

	fmt.Println("Program finished successfully!")
}

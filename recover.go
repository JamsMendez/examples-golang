package main

import "fmt"

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")

		panic("highlow() low greater than high")
	}

	defer fmt.Printf("Deferred highlow(%d, %d)\n", high, low)

	fmt.Printf("Call: highlow(%d, %d)\n", high, low)

	highlow(high, low+1)
}

func main() {
	defer func() {
		handler := recover()
		if handler != nil {
		  // Allow you to regain control after a panic
		  // backoup, write logs, etc ...
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)

	fmt.Println("Program finished successfully!")
}

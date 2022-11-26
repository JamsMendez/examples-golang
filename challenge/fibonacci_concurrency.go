package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type Result struct {
	Index int
	Value float64
}

func fib(number float64) float64 {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	return x
}

const cmdQuit = "quit"

func main() {
	start := time.Now()

	result := make(chan Result)
	inputUser := make(chan string)

	go func() {
		for {
			var input string
			fmt.Scanf("%s ...", &input)

			inputUser <- input

			if input == cmdQuit {
				break
			}
		}
	}()

	i := 0
	running := true

	for running {
		select {
		case value := <-inputUser:
			if value == cmdQuit {
				close(inputUser)
				close(result)
				running = false
			}

			if value == "" {
				i += 1
				go func(index int) {
					n := fib(float64(index))
					result <- Result{Index: index, Value: n}
				}(i)
			}

		case res := <-result:
			fmt.Printf("Fib(%v): %v\n", res.Index, res.Value)
		}

		time.Sleep(time.Millisecond * 500)
		fmt.Println("...")
	}

	fmt.Println(runtime.NumGoroutine())

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func mainConcurrency() {
	start := time.Now()

	result := make(chan Result, 14)
	var numberResults int

	for i := 1; i < 15; i++ {
		go func(index int) {
			n := fib(float64(index))
			result <- Result{Index: index, Value: n}
		}(i)
	}

	for res := range result {
		numberResults += 1
		fmt.Printf("Fib(%v): %v\n", res.Index, res.Value)

		if numberResults == 14 {
			close(result)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

/*

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func fib(number float64, ch chan string) {
    x, y := 1.0, 1.0
    for i := 0; i < int(number); i++ {
        x, y = y, x+y
    }

    r := rand.Intn(3)
    time.Sleep(time.Duration(r) * time.Second)

    ch <- fmt.Sprintf("Fib(%v): %v\n", number, x)
}

func main() {
    start := time.Now()

    size := 15
    ch := make(chan string, size)

    for i := 0; i < size; i++ {
        go fib(float64(i), ch)
    }

    for i := 0; i < size; i++ {
        fmt.Printf(<-ch)
    }

    elapsed := time.Since(start)
    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

package main

import (
    "fmt"
    "time"
)

var quit = make(chan bool)

func fib(c chan int) {
    x, y := 1, 1

    for {
        select {
            case c <- x:
                x, y = y, x+y
            case <-quit:
                fmt.Println("Done calculating Fibonacci!")
            return
        }
    }
}

func main() {
    start := time.Now()

    command := ""
    data := make(chan int)

    go fib(data)

    for {
        num := <-data
        fmt.Println(num)
        fmt.Scanf("%s", &command)
        if command == "quit" {
            quit <- true
            break
        }
    }

    time.Sleep(1 * time.Second)

    elapsed := time.Since(start)
    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

*/

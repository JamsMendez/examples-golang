package main

import (
	"context"
	"fmt"
	"time"
)

// - Avoid Shared State: Minimize the use of shared mutable state between
// goroutines to reduce the risk of race conditions and data races.
// - Use Channels for Communication: Prefer channels for communication between
// goroutines over shared memory access, as channels enforce synchronization
// and facilitate safe communication.
// - Avoid Goroutine Leaks: Ensure that goroutines are properly managed and cleaned
// up to avoid leaks that could lead to resource exhaustion.
// - Graceful Shutdown: Implement mechanisms for graceful shutdown of concurrent
// operations to prevent resource leaks and ensure clean program termination.
// - Use `context` Package for Cancellation: Utilize the `context` package for
// managing timeouts, deadlines, and cancellation of long-running operations
// across goroutines.

func mainMultiplexing() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go producerMultiplexing("Channel 1", ch1)
	go producerMultiplexing("Channel 2", ch2)

	for i := 0; i < 5; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received from Channel 1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from Channel 2:", msg)
		}
	}
}

func producerMultiplexing(name string, ch chan<- string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- fmt.Sprintf("%s - Message %d", name, i)
	}
	close(ch)
}

func mainControllingCapacity() {
	ch := make(chan int, 3) // Buffered channel with capacity 3

	go producer(ch)

	for i := 0; i < 5; i++ {
		fmt.Println("Consuming:", <-ch)
		time.Sleep(500 * time.Millisecond)
	}
}

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Producing:", i)
		ch <- i // Send data to the buffered channel
	}
	close(ch) // Close the channel after sending all data
}

func mainContextManage() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go longRunningTask(ctx)

	// Wait for goroutine to finish or timeout
	v1 := <-ctx.Done()
	fmt.Println("v1 ", v1)
	v2 := <-ctx.Done()
	fmt.Println("v3 ", v2)
	v3 := <-ctx.Done()
	fmt.Println("v3 ", v3)

	fmt.Println("Operation canceled:", ctx.Err())

	// <-time.After(5 * time.Second)
}

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed successfully")
	case v := <-ctx.Done():
		fmt.Println("Task canceled:", ctx.Err(), v)
	}
}

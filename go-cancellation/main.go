package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	fileRoot = "root.txt"
)

func finder(ctx context.Context, ID int, result chan<- string) {
	done := make(chan struct{})

	var file string

	go func() {
		workTime := rand.Intn(10) + 1

		if ID == 3 {
			workTime = 4
			file = fileRoot
		}

		duration := time.Duration(workTime) * time.Second
		fmt.Printf("Finder %d work time %s\n", ID, duration.String())
		<-time.After(duration)
		close(done)
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("Finder %d canceled\n", ID)
		close(done)

	case <-done:
		fmt.Printf("Finder %d finished\n", ID)
		result <- fmt.Sprintf("Finder %d: %s", ID, file)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	result := make(chan string)

	for i := 0; i < 5; i++ {
		ID := i + 1
		go finder(ctx, ID, result)
	}

	go func() {
		for v := range result {
			fmt.Println(v)
			if v == fileRoot {
				cancel()
			}
		}
	}()

	<-ctx.Done()

	<-time.After(1 * time.Second)

	close(result)
}

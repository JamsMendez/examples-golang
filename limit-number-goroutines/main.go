package main

import (
	"log"
	"sync"
	"time"
)

const maxNum = 3

func main() {
	runLimit()
}

func runLimit() {
	var wg sync.WaitGroup

	limiter := make(chan int, 3)

	for i := 0; i < 5; i++ {
		id := i + 1

		wg.Add(1)

		go func() {
			defer func() {
				wg.Done()
				<-limiter
			}()

			limiter <- id

			delay := time.Millisecond * time.Duration(100*(id))
			log.Printf(" Start #%d (delay %v)\n", id, delay)
			time.Sleep(delay)

			log.Printf(" End #%d (delay %v)\n", id, delay)

		}()
	}

	wg.Wait()

	log.Println("Finish")
}

func runWithLimit() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		id := i + 1
		wg.Add(1)
		go func() {
			defer wg.Done()

			delay := time.Millisecond * time.Duration(100*(id))
			log.Printf(" Start #%d (delay %v)\n", id, delay)
			time.Sleep(delay)

			log.Printf(" End #%d (delay %v)\n", id, delay)
		}()
	}

	wg.Wait()

	log.Println("Finish")
}

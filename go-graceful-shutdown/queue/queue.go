package queue

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Run() {
	var wg sync.WaitGroup

	messageQueue := make(chan string, 10)
	shutdown := make(chan struct{})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	numSubscribers := 3

	for i := 0; i < numSubscribers; i++ {
		wg.Add(1)

		go subscriber(i, &wg, messageQueue)
	}

	go func() {
		j := 1

		var queueStop bool

		for {
			select {
			case <-shutdown:
				queueStop = true
			default:
				messageQueue <- fmt.Sprintf("Message %d", j)

				time.Sleep(time.Millisecond * 500)

				j = j + 1
			}

			if queueStop {
				break
			}
		}

		close(messageQueue)
	}()

	<-stop
	fmt.Printf("\nShutting down gracefully ...")

	wg.Wait()

	fmt.Println("All subscrubers have finished. Exiting ...")
}

func subscriber(ID int, wg *sync.WaitGroup, messages chan string) {
	defer wg.Done()

	for message := range messages {
		fmt.Printf("Subscriber %d received: %s\n", ID, message)
		time.Sleep(time.Millisecond * 500)
	}

	fmt.Printf("Subscriber %d shutting down.\n", ID)
}

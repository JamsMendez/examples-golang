package queue

import (
	"fmt"
	"sync"
	"time"
)

func producer(queue *Queue, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		item := id*10 + i
		queue.Enqueue(item)
		<-time.After(5 * time.Second)
	}
}

func consumer(queue *Queue, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		item := queue.Dequeue()
		fmt.Printf("Consumer %d consumed: %d\n", id, item)
		<-time.After(5 * time.Second)
	}
}

func Run() {
	queue := NewQueue()

	var wg sync.WaitGroup

	// producers
	// for i := 0; i < 3; i++ {
	wg.Add(1)
	go producer(queue, 1, &wg)
	// }

	// consumers
	// for i := 0; i < 3; i++ {
	wg.Add(1)
	go consumer(queue, 1, &wg)
	// }

	wg.Wait()

	fmt.Println("All producers and consumers have finished")
}

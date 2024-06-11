package worker

import (
	"fmt"
	"sync"
	"time"
)

func WorkerFunc(id int, startCond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()

	startCond.L.Lock()
	fmt.Printf("Worker %d pre wait ...\n", id)
	startCond.Wait()
	fmt.Printf("Worker %d post wait ...\n", id)
	startCond.L.Unlock()

	fmt.Printf("Worker %d started working\n", id)
	<-time.After(5 * time.Second)
	fmt.Printf("Worker %d finished working\n", id)
}

func Run() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	startCond := sync.NewCond(&mu)

	numWorkers := 5

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go WorkerFunc(i, startCond, &wg)
	}

	<-time.After(3 * time.Second)
	fmt.Println("Signaling workers to start")

	startCond.Broadcast()

	wg.Wait()

	fmt.Println("All workers have finished")
}

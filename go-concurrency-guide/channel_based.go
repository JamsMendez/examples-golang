package main

import (
	"fmt"
	"sync"
	"time"
)

type DataChannelBased struct {
	Value string
}

func mainChannleBasedComm() {
	data := make(chan *DataChannelBased)

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go modifyData(data, &wg)
	}

	// Receive modified data from channels
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Received:", <-data)
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("All Goroutines Done!")
}

func modifyData(data chan<- *DataChannelBased, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate some work
	time.Sleep(200 * time.Millisecond)

	// Modify and send data through the channel
	data <- &DataChannelBased{Value: "modified"}
	fmt.Println("Data modified by goroutine")
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func RunSync() {
	nReq := 10

	var reqsTime time.Duration
	var slowest time.Duration

	semLen := 3
	semaphoreChan := make(chan struct{}, semLen)
	durationChan := make(chan time.Duration)

	receiveCount := 0
	sendCount := 0

	start := time.Now()

	for receiveCount < nReq {
		select {
		case semaphoreChan <- struct{}{}:
			if sendCount < nReq {
				sendCount++
				go requestSync(durationChan, semaphoreChan)
			}

		case reqDuration := <-durationChan:
			fmt.Printf("single requesttime taken: %v\n", reqDuration)
			receiveCount++
			reqsTime += reqDuration
			if reqDuration > slowest {
				slowest = reqDuration
			}
		}
	}

	close(semaphoreChan)
	close(durationChan)

	fmt.Printf("Time taken to serve all requests: %v\n", time.Since(start))
	fmt.Printf("Slowest request time taken: %v\n", slowest)
	fmt.Printf("Requests total computation time: %v\n", reqsTime)
}

func requestSync(d chan time.Duration, sem chan struct{}) {
	defer func() {
		<-sem
	}()

	start := time.Now()
	_, _ = http.Get("https://httpbin.org/get")

	d <- time.Since(start)
}

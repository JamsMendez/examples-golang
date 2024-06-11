package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"go-semaphore-channel-pattern/semaphore"
)

func RunBasic() {
	nReq := 10
	var wg sync.WaitGroup
	wg.Add(nReq)

	semLen := 3
	semaphore := make(chan struct{}, semLen)

	start := time.Now()

	for i := 0; i < nReq; i++ {
		semaphore <- struct{}{}
		go requestBasic(&wg, semaphore)
	}

	wg.Wait()

	fmt.Printf("time take to server all requests: %v\n", time.Since(start))
}

func requestBasic(wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()

	defer func() {
		<-semaphore
	}()

	start := time.Now()
	_, _ = http.Get("https://httpbin.org/get")

	fmt.Printf("sigle request time taken: %v\n", time.Since(start))
}

func RunBasicWithSemaphore() {
	nReq := 10

	var wg sync.WaitGroup
	wg.Add(nReq)

	tickets, timeout := 3, 3*time.Second
	s := semaphore.New(tickets, timeout)

	start := time.Now()
	for i := 0; i < nReq; i++ {
		if err := s.Acquire(); err != nil {
			log.Fatal("acquire: ", err)
		}

		go requestBasicWithSemaphore(&wg, s)
	}

	wg.Wait()

	fmt.Printf("time take to server all requests: %v\n", time.Since(start))
}

func requestBasicWithSemaphore(wg *sync.WaitGroup, s semaphore.Semaphore) {
	defer wg.Done()

	defer func() {
		if err := s.Release(); err != nil {
			log.Fatal("release: ", err)
		}
	}()

	start := time.Now()
	_, _ = http.Get("https://httpbin.org/get")

	fmt.Printf("sigle request time taken: %v\n", time.Since(start))
}

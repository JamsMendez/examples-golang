package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"go-semaphore-channel-pattern/semaphore"
)

func main() {
	RunBasic()
	<-time.After(5 * time.Second)
	fmt.Println()
	RunBasicWithSemaphore()
}

func RunSemaphore() {
	tickets, timeout := 1, 3*time.Second

	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		log.Fatal("acquire: ", err)
	}

	if err := s.Release(); err != nil {
		log.Fatal("release: ", err)
	}

	fmt.Println("finish semaphore")
}

func RunSemaphoreWithoutTimeouts() {
	tickets, timeout := 0, 0*time.Second

	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		if err != semaphore.ErrNoTickets {
			log.Fatal("acquire: ", err)
		}
	}

	// No tickets left, can't work :/
	os.Exit(1)
}

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /sleep", getSleep)
	mux.HandleFunc("GET /sleep/{n}", getSleepN)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}

type SleepResponseSchema struct {
	Time string `json:"timeTaken"`
}

func getSleep(w http.ResponseWriter, r *http.Request) {
	sleepDuration := time.Millisecond * time.Duration(rand.Int64N(1000))
	<-time.After(sleepDuration)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	sleep := SleepResponseSchema{Time: sleepDuration.String()}
	_ = json.NewEncoder(w).Encode(sleep)
}

func makeRequest(d chan<- time.Duration, semaphore chan struct{}) {
	defer func() {
		<-semaphore
	}()

	resp, _ := http.Get("http://localhost:8080/sleep")

	sleep := SleepResponseSchema{}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("reponse body close: ", err)
		}
	}()

	_ = json.NewDecoder(resp.Body).Decode(&sleep)
	reqDuration, _ := time.ParseDuration(sleep.Time)
	d <- reqDuration
}

func getSleepN(w http.ResponseWriter, r *http.Request) {
	log.Println("getSleepN...")

	t := time.Now()

	var reqsTime time.Duration
	var slowest time.Duration

	d := make(chan time.Duration)
	semaphore := make(chan struct{}, 100)

	n, _ := strconv.Atoi(r.PathValue("n"))

	receiveCount := 0
	sendCount := 0

	for receiveCount < n {
		log.Printf("activeGoroutine: %d | sendCount: %d | receiveCount: %d", runtime.NumGoroutine(), sendCount, receiveCount)

		select {
		case semaphore <- struct{}{}:
			if sendCount < n {
				sendCount++
				go makeRequest(d, semaphore)
			}
		case reqDuration := <-d:
			receiveCount++
			reqsTime += reqDuration
			if reqDuration > slowest {
				slowest = reqDuration
			}
		}
	}

	close(semaphore)
	close(d)

	_, _ = fmt.Fprintf(w, "Benchmark time taken: %v", time.Since(t))
	_, _ = fmt.Fprintf(w, "\nSlowest request time taken: %v", slowest)
	_, _ = fmt.Fprintf(w, "\nRequests total computation time: %v", reqsTime)

	log.Printf("activeGoroutine:%d | sendCount:%d | receiveCount:%d", runtime.NumGoroutine(), sendCount, receiveCount)
}

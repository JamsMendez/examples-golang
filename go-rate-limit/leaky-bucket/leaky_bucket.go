package leakybucket

import (
	"log"
	"net/http"
	"time"
)

type RateLimiter interface {
	Allow() bool
}

type LeakyBucket struct {
	queue    chan (chan struct{})
	capacity int
	rps      int
}

func (lb *LeakyBucket) Allow() bool {
	ch := make(chan struct{})
	if !lb.AddToQueue(ch) {
		return false
	}

	<-ch
	return true
}

func (lb *LeakyBucket) AddToQueue(ch chan struct{}) bool {
	if len(lb.queue) < lb.capacity {
		lb.queue <- ch
		return true
	}

	return false
}

func NewLeakyBucket(count, duration int) RateLimiter {
	lb := &LeakyBucket{
		make(chan chan struct{}, count),
		count,
		duration,
	}

	go AddTicker(lb)

	return lb
}

func AddTicker(lb *LeakyBucket) {
	for value := range lb.queue {
		value <- struct{}{}
		k := int64(1000000 / lb.rps)
		if k < 1 {
			k = 1
		}
		<-time.After(time.Duration(k) * time.Microsecond)
	}
}

func RunRateLimiter() {
	rl := NewLeakyBucket(10, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if rl.Allow() {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusTooManyRequests)
	})

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

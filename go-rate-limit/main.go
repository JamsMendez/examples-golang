package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"golang.org/x/time/rate"

	rateLimiter "go-rate-limit/http-request"
)

func main() {
	runTickRateLimit()
}

func runMain() {
	// runRateLimitWaitMultiple()

	limiter := rate.NewLimiter(rate.Every(5*time.Second), 5)
	// limiter := rate.NewLimiter(Per(5, 5*time.Second), 5)

	for i := 0; i < 20; i++ {
		err := limiter.Wait(context.Background())
		if err != nil {
			fmt.Println(err)

			continue
		}

		fmt.Println(i, " success ", time.Now())
	}
}

func runSimpleRateLimit() {
	requests := make(chan int, 5)

	// simulation of request o async process
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	// limiter is a channel will receive a value every second.
	// This is the regulation in our rate limiting scheme.
	limiter := time.Tick(1000 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("Request 1: ", req, time.Now())
	}

	burstLimiter := make(chan time.Time, 3)

	// if isn't fill up, the first one will take 3 seconds
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}

	close(burstRequests)

	for req := range burstRequests {
		<-burstLimiter
		fmt.Println("Request 2:", req, time.Now())
	}
}

func runTickRateLimit() {
	doWork := func(taskID int, limiter <-chan time.Time) {
		<-limiter
		// Perform task operations
		fmt.Println("TaskID: ", taskID)
	}

	tasks := make(chan int, 10)
	for i := 0; i < 10; i++ {
		tasks <- i
	}
	fmt.Println("Set tasks ...")

	limiter := time.Tick(time.Second * 2)

	fmt.Println("Get tasks ...")

	for taskID := range tasks {
		go doWork(taskID, limiter)
		fmt.Println(taskID, len(tasks), cap(tasks))
		if len(tasks) == 0 {
			close(tasks)
		}
	}

	fmt.Println("Wait ... ")
	<-time.After(25 * time.Second)
	fmt.Println("Finish...")
}

// rate limiter in server http and client http
func runServerRateLimit() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	fmt.Println("Server listening")

	go func() {
		time.Sleep(2 * time.Second)

		rateLimiter.RunClient()
	}()

	if err := http.ListenAndServe(":3000", limitMiddleware(mux)); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}

// handle simple response
func okHandler(w http.ResponseWriter, r *http.Request) {
	buffer := []byte("response success")
	w.Write(buffer)
}

// middleware with rate limit
func limitMiddleware(next http.Handler) http.Handler {
	limiter := rateLimiter.NewIPRateLimiter(rate.Every(1*time.Second), 20)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mLimiter := limiter.GetLimiter(r.RemoteAddr)

		if !mLimiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func runRateLimitWaitMultiple() {
	apiConnection := OpenMulti()

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			v, err := apiConnection.Read(context.Background())
			if err != nil {
				fmt.Printf("%v Get Error: %v\n", time.Now().Format("15:04:05"), err)
				return
			}

			fmt.Printf("%v %v\n", time.Now().Format("15:04:05"), v)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			err := apiConnection.Resolve(context.Background())
			if err != nil {
				fmt.Printf("%v Resolve Error: %v\n", time.Now().Format("15:04:05"), err)
				return
			}

			fmt.Printf("%v Resolved\n", time.Now().Format("15:04:05"))
		}()
	}

	wg.Wait()
}

type ApiConnMulti struct {
	apiLimit,
	dbLimit RateLimiter
}

func (api *ApiConnMulti) Read(ctx context.Context) (string, error) {
	if err := api.apiLimit.Wait(ctx); err != nil {
		return "", nil
	}

	// Do work
	return "Read", nil
}

func (api *ApiConnMulti) Resolve(ctx context.Context) error {
	if err := api.dbLimit.Wait(ctx); err != nil {
		return err
	}

	// Do work
	return nil
}

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

type multiLimiter struct {
	limiters []RateLimiter
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}

	sort.Slice(limiters, byLimit)

	return &multiLimiter{limiters: limiters}
}

func (ml *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range ml.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (l *multiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}

func OpenMulti() *ApiConnMulti {
	return &ApiConnMulti{
		apiLimit: MultiLimiter(
			rate.NewLimiter(Per(2, time.Second), 1),
			rate.NewLimiter(Per(5, time.Second), 5),
		),
		dbLimit: MultiLimiter(
			rate.NewLimiter(rate.Every(time.Second), 1),
		),
	}
}

func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

func runRateLimitWait() {
	defer fmt.Println("exiting application")

	apiConnection := Open()

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			v, err := apiConnection.Read(context.Background())
			if err != nil {
				fmt.Printf("%v Get Error: %v\n", time.Now().Format("15:04:05"), err)
				return
			}

			fmt.Printf("%v %v\n", time.Now().Format("15:04:05"), v)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			err := apiConnection.Resolve(context.Background())
			if err != nil {
				fmt.Printf("%v Resolve Error: %v\n", time.Now().Format("15:04:05"), err)
				return
			}

			fmt.Printf("%v Resolved\n", time.Now().Format("15:04:05"))
		}()
	}

	wg.Wait()
}

type ApiConn struct {
	rateLimiter *rate.Limiter
}

func Open() *ApiConn {
	return &ApiConn{
		rateLimiter: rate.NewLimiter(rate.Every(time.Second), 1),
	}
}

func (api *ApiConn) Read(ctx context.Context) (string, error) {
	if err := api.rateLimiter.Wait(ctx); err != nil {
		return "", nil
	}

	// Do work
	return "Read", nil
}

func (api *ApiConn) Resolve(ctx context.Context) error {
	if err := api.rateLimiter.Wait(ctx); err != nil {
		return err
	}

	// Do work
	return nil
}

func runRateLimitDrip() {
	b := Bucket{
		Capacity:     60,
		DripInterval: 3 * time.Second,
		PerDrip:      10,
	}

	b.Start()
	defer b.Stop()

	count := 0
	for {
		err := b.Consume(2)
		if err != nil {
			fmt.Println("Sleeping a bit")
			time.Sleep(time.Second)
		} else {
			count++
			fmt.Println("Consuming 2")
		}

		if count > 100 {
			break
		}
	}
}

type Bucket struct {
	Capacity     int
	DripInterval time.Duration
	PerDrip      int
	consumed     int
	started      bool
	kill         chan bool
	m            sync.Mutex
}

func (b *Bucket) Start() error {
	if b.started {
		return errors.New("bucket was already started")
	}

	ticker := time.NewTicker(b.DripInterval)
	b.started = true
	b.kill = make(chan bool, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				b.m.Lock()
				b.consumed -= b.PerDrip
				fmt.Println("Start.b.consumed: ", b.consumed, " - ", b.PerDrip, " = ", b.consumed)

				if b.consumed < 0 {
					b.consumed = 0
					fmt.Println("Start.b.command is 0")
				}

				b.m.Unlock()
			case <-b.kill:
				return
			}
		}
	}()

	return nil
}

func (b *Bucket) Stop() error {
	if !b.started {
		return errors.New("bucket was never started")
	}

	b.kill <- true

	return nil
}

func (b *Bucket) Consume(amt int) error {
	b.m.Lock()
	defer b.m.Unlock()

	if b.Capacity-b.consumed < amt {
		return errors.New("not enough capcity")
	}

	b.consumed += amt

	fmt.Println("Consume.b.consumed: ", b.consumed)

	return nil
}

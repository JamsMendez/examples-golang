package main

import (
	"fmt"
	"time"
)

func main() {
	result1 := AsyncInvoker(fakeTask(1 * time.Second))
	result2 := AsyncInvoker(fakeTask(2 * time.Second))
	result3 := AsyncInvoker(fakeTask(3 * time.Second))

	for i := 0; i < 5; i++ {
		<-time.After(1 * time.Second)
		fmt.Println("working on other things")
		if result2.IsDone() {
			fmt.Println(result2.Get())
		}

	}

	fmt.Println(result1.Get())
	fmt.Println(result3.Get())
}

// AsyncFunction
type AsyncTask func() (string, error)

// Future
type AsyncResult struct {
	result string
	err    error
	done   chan struct{}
}

func (ar *AsyncResult) Get() (string, error) {
	<-ar.done
	return ar.result, ar.err
}

func (ar *AsyncResult) IsDone() bool {
	select {
	case <-ar.done:
		return true
	default:
		return false
	}
}

func NewAsyncResult() *AsyncResult {
	return &AsyncResult{
		done: make(chan struct{}),
	}
}

// Implementacion de la AsyncFunction
func fakeTask(duration time.Duration) AsyncTask {
	return func() (string, error) {
		<-time.After(duration)
		// fmt.Println("task ", duration.String())
		return "Task complete", nil
	}
}

// Client
func AsyncInvoker(task AsyncTask) *AsyncResult {
	result := NewAsyncResult()

	go func() {
		r, err := task()
		result.result = r
		result.err = err

		close(result.done)
	}()

	return result
}

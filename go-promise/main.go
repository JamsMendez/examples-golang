package main

import (
	"fmt"
	"time"
)

type Promise struct {
	result chan any
	err    chan error
}

func NewPromise(fn func() (any, error)) *Promise {
	p := &Promise{
		result: make(chan any),
		err:    make(chan error),
	}

	go func() {
		result, err := fn()
		if err != nil {
			p.err <- err
		} else {
			p.result <- result
		}
	}()

	return p
}

func (p *Promise) Then(resolve func(any), reject func(error), finally func()) {
	select {
	case result := <-p.result:
		resolve(result)
	case err := <-p.err:
		reject(err)
	}
	close(p.result)
	close(p.err)

	finally()
}

func main() {
	fetch := func() (any, error) {
		<-time.After(2 * time.Second)
		return "JSON", nil
	}

	promise := NewPromise(fetch)

	done := make(chan struct{})

	go promise.Then(
		func(value any) {
			fmt.Println("Success: ", value)
		},
		func(err error) {
			fmt.Println("Error: ", err)
		},
		func() {
			fmt.Println("Finish")
			close(done)
		},
	)

	<-done
}

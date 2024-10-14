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
			close(p.result)
		} else {
			p.result <- result
			close(p.err)
		}
	}()

	return p
}

func (p *Promise) Then(resolve func(any), reject func(error), finally func()) {
	select {
	case result := <-p.result:
		resolve(result)
	case <-p.err:
		if reject != nil {
			reject(nil)
		}
	}
	finally()
}

func main() {
	fetch := func() (any, error) {
		<-time.After(2 * time.Second)
		return "JSON", nil
	}

	promise := NewPromise(fetch)

	done := make(chan struct{})

	promise.Then(
		func(value any) {
			fmt.Println("Response: ", value)
		},
		func(err error) {
			fmt.Println("Error: ", err)
		},
		func() {
			close(done)
		},
	)

	<-done
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	account := New(100)
	_, _ = account.Deposit(200)
	_, _ = account.Withdraw(100)

	fmt.Println(account.Balance())

	fmt.Println("Async")

	var wg sync.WaitGroup

	for i := 0; i < 10_000; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			if index%2 == 0 {
				_, _ = account.Withdraw(1)
			} else {
				_, _ = account.Deposit(1)
			}
		}(i)
	}

	wg.Wait()

	// 200
	fmt.Println(account.Balance())
}

type BankAccount interface {
	Balance() (int, error)
	Deposit(amount int) (int, error)
	Withdraw(amount int) (int, error)
	Close()
}

type Request interface {
	Call()
}

type scheduler struct {
	queue chan Request
}

func newScheduler(size int) *scheduler {
	s := &scheduler{
		queue: make(chan Request, size),
	}

	go s.Dispacher()

	return s
}

func (s *scheduler) Dispacher() {
	for req := range s.queue {
		req.Call()
	}
}

func (s *scheduler) Enqueue(req Request) {
	s.queue <- req
}

func (s *scheduler) Shutdown() {
	close(s.queue)
}

type result struct {
	value int
	err   error
	done  chan struct{}
}

type txRequest struct {
	action func() (int, error)
	result *result
}

func (r *txRequest) Call() {
	value, err := r.action()
	r.result.value = value
	r.result.err = err
	close(r.result.done)
}

type bankAccount struct {
	balance   int
	scheduler *scheduler
}

func New(balance int) *bankAccount {
	account := &bankAccount{
		balance:   balance,
		scheduler: newScheduler(10),
	}

	return account
}

func (b *bankAccount) Balance() (int, error) {
	res := &result{
		done: make(chan struct{}),
	}

	req := &txRequest{
		action: func() (int, error) {
			return b.balance, nil
		},
		result: res,
	}

	b.scheduler.Enqueue(req)

	<-res.done

	return res.value, res.err
}

func (b *bankAccount) Deposit(amount int) (int, error) {
	res := &result{
		done: make(chan struct{}),
	}

	req := &txRequest{
		action: func() (int, error) {
			b.balance += amount

			return b.balance, nil
		},
		result: res,
	}

	b.scheduler.Enqueue(req)

	<-res.done

	return res.value, res.err
}

func (b *bankAccount) Withdraw(amount int) (int, error) {
	res := &result{
		done: make(chan struct{}),
	}

	req := &txRequest{
		action: func() (int, error) {
			if b.balance >= amount {
				b.balance -= amount

				return b.balance, nil
			}

			return b.balance, fmt.Errorf("balance amount insufficient")
		},
		result: res,
	}

	b.scheduler.Enqueue(req)

	<-res.done

	return res.value, res.err
}

func (b *bankAccount) Close() {
	b.scheduler.Shutdown()
}

package semaphore

import (
	"errors"
	"time"
)

var (
	ErrNoTickets     = errors.New("semaphore: could not aquire semaphore")
	ErrIlegalRelease = errors.New("smaphore: can't release the semaphore without acquiring it first")
)

type Semaphore interface {
	Acquire() error
	Release() error
}

type TicketSemaphore struct {
	sem     chan struct{}
	timeout time.Duration
}

func (t *TicketSemaphore) Acquire() error {
	select {
	case t.sem <- struct{}{}:
		return nil
	case <-time.After(t.timeout):
		return ErrNoTickets
	}
}

func (t *TicketSemaphore) Release() error {
	select {
	case <-t.sem:
		return nil
	case <-time.After(t.timeout):
		return ErrIlegalRelease
	}
}

func New(tickets int, timeout time.Duration) Semaphore {
	return &TicketSemaphore{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}

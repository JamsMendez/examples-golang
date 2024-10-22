package jobqueue

import (
	"log"
	"sync"
	"sync/atomic"
)

// Job defines the inteface for job that can be submitted to the job queue
type Job interface {
	Execute()
}

// JobQueue represents a queue that processes jobs in FIFO order
type JoeQueue struct {
	queue  chan Job
	wg     sync.WaitGroup
	quit   chan struct{}
	closed atomic.Bool
}

// Enqueue adds a job to the queue
func (jq *JoeQueue) Enqueue(job Job) {
	if jq.closed.Load() {
		log.Println("Attemp to enqueue on closed queue")
		return
	}

	// jq.queue <- job

	select {
	case jq.queue <- job:
		// Job enqueued successfully
	default:
		log.Println("queue if full. Job rejected")
	}
}

// worker processes jobs as they arrive in the queue
func (jq *JoeQueue) Worker() {
	defer jq.wg.Done()

	for {
		select {
		case job := <-jq.queue:
			job.Execute()
		case <-jq.quit:
			return
		}
	}
}

// Close gracefully shuts down the job queue after all jobs are processed
func (jq *JoeQueue) Close() {
	if jq.closed.CompareAndSwap(false, true) {
		close(jq.quit)
		jq.wg.Wait()

		close(jq.queue)
	}
}

func New(bufferSize int) *JoeQueue {
	jq := &JoeQueue{
		queue: make(chan Job, bufferSize),
		quit:  make(chan struct{}),
	}

	jq.wg.Add(1)
	go jq.Worker()

	return jq
}

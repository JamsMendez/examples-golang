package jobqueue

import (
	"sync/atomic"
	"testing"
	"time"
)

// Define a simple Job implementation for testing
type TestJob struct {
	id       int
	executed atomic.Bool
}

func (t *TestJob) Execute() {
	t.executed.Store(true)
}

// TestJobQueue ensures that jobs are processed in the order they are added
func TestJobQueue(t *testing.T) {
	queue := New(3)
	defer queue.Close()

	jobs := []*TestJob{
		{id: 1},
		{id: 2},
		{id: 3},
		{id: 4},
	}

	for _, job := range jobs {
		queue.Enqueue(job)
	}

	<-time.After(100 * time.Millisecond)

	for i := 0; i < 3; i++ {
		if !jobs[i].executed.Load() {
			t.Errorf("expected job %d was executed to true, got %v", jobs[i].id, jobs[i].executed.Load())
		}
	}

	if jobs[3].executed.Load() {
		t.Errorf("expected job 4 was executed to false, got %v", jobs[3].executed.Load())
	}
}

// TestJobQueueClosure test that the queue correctly stops accepting jobs after closure.
func TestJobQueueClosure(t *testing.T) {
	queue := New(1)
	defer queue.Close()

	job := &TestJob{id: 1}

	queue.Enqueue(job)
	queue.Close() // Ensure no more jobs can be enqueued and the queue shuts down

	// Allow time for job to be processed and worker to exit
	<-time.After(100 * time.Millisecond)

	if !job.executed.Load() {
		t.Errorf("expected job was executed to true, got %v", job.executed.Load())
	}

	// Test that no more jobs can be added after closure
	done := make(chan struct{})
	go func() {
		queue.Enqueue(&TestJob{id: 2})
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Errorf("enqueue after Close(), did not return promptly")
	}
}

package eventqueue

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestNewEventQueue(t *testing.T) {
	queueName := "test_queue"
	bufferSize := 1

	queue := NewEventQueue(queueName, bufferSize)

	if queue.close == nil {
		t.Fatalf("expected chan close != nil, got %v", queue.close)
	}

	if queue.queue == nil {
		t.Fatalf("expected chan events != nil, got %v", queue.queue)
	}

	if queue.drain == nil {
		t.Fatalf("expected chan drain != nil, got %v", queue.drain)
	}

	if queue.name != queueName {
		t.Fatalf("expected queue name %s, got %v", queueName, queue.name)
	}

	capacity := cap(queue.queue)
	if capacity != bufferSize {
		t.Fatalf("expected queue size %d, got %v", bufferSize, capacity)
	}
}

func TestNilEventQueue(t *testing.T) {
	var queue *EventQueue
	queue.Stop()

	if queue != nil {
		t.Fatalf("expected queue to be nil, got %v", queue)
	}
}

func TestStopWithoutRun(t *testing.T) {
	queue := NewEventQueue("", 1)
	queue.Stop()
}

func TestCloseEventQueueMultipleTimes(t *testing.T) {
	queue := NewEventQueue("", 1)
	queue.Stop()
	// closing event queue twice should not cause panic
	queue.Stop()
	queue.Stop()
}

func TestDrained(t *testing.T) {
	queue := NewEventQueue("", 1)
	queue.Run()

	// Stopping queue should drain it as well
	queue.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	select {
	case <-queue.close:
	case <-ctx.Done():
		t.Log("timed out waiting for queue to be drained")
		t.Fatal()
	}
}

func TestNilEvent(t *testing.T) {
	queue := NewEventQueue("", 1)
	res, err := queue.Enqueue(nil)

	if res != nil {
		t.Fatalf("expected res to be nil, got %v", res)
	}

	if !errors.Is(err, ErrUnableToEnqueueEvent) {
		t.Fatalf("expected error %v, got %v", ErrUnableToEnqueueEvent, err)
	}
}

func TestNewEvent(t *testing.T) {
	event := NewEvent(&FakeEvent{})
	if event.Metadata == nil {
		t.Fatalf("expected event.Metadata != nil, got %v", event.Metadata)
	}

	if event.eventResults == nil {
		t.Fatalf("expected event.eventResults != nil, got %v", event.eventResults)
	}

	if event.cancelled == nil {
		t.Fatalf("expected event.cancelled != nil, got %v", event.cancelled)
	}
}

type FakeEvent struct{}

func (d *FakeEvent) Handle(value chan any) {
	value <- struct{}{}
}

func TestEventCancelAfterQueueClosed(t *testing.T) {
	queue := NewEventQueue("", 1)
	queue.Run()

	event := NewEvent(&FakeEvent{})
	_, err := queue.Enqueue(event)
	if err != nil {
		t.Fatalf("expected err == nil, got %v", err)
	}

	// Event should not have been cancelled since queue was not closed
	if event.WasCancelled() {
		t.Fatalf("expected event.WasCancelled false, got true")
	}

	queue.Stop()

	event = NewEvent(&FakeEvent{})
	_, err = queue.Enqueue(event)
	if err != nil {
		t.Fatalf("expected err == nil, got %v", err)
	}

	if !event.WasCancelled() {
		t.Fatalf("expected event.WasCancelled true, got false")
	}
}

type NewHangEvent struct {
	Channel   chan struct{}
	name      string
	processed bool
}

func (n *NewHangEvent) Handle(value chan any) {
	// blocked ...
	<-n.Channel
	n.processed = true
	value <- struct{}{}
}

func CreateHangEvent(name string) *NewHangEvent {
	return &NewHangEvent{
		Channel: make(chan struct{}),
		name:    name,
	}
}

func TestDrain(t *testing.T) {
	queue := NewEventQueue("", 1)
	queue.Run()

	nh1 := CreateHangEvent("nh1")
	nh2 := CreateHangEvent("nh2")
	nh3 := CreateHangEvent("nh3")

	event := NewEvent(nh1)
	_, err := queue.Enqueue(event)
	if err != nil {
		t.Fatalf("expected err == nil, got %v", err)
	}

	event2 := NewEvent(nh2)
	event3 := NewEvent(nh3)

	_, err = queue.Enqueue(event2)
	if err != nil {
		t.Fatalf("expected err == nil, got %v", err)
	}

	var (
		rcvChan <-chan any
		err3    error
	)

	enq := make(chan struct{})

	go func() {
		rcvChan, err3 = queue.Enqueue(event3)
		if err3 != nil {
			t.Errorf("expected err to be nil, got %v", err3)
		}
		enq <- struct{}{}
	}()

	close(nh1.Channel)

	<-enq

	go queue.Stop()

	<-queue.drain

	close(nh2.Channel)

	_, ok := <-rcvChan
	if ok {
		t.Fatal("expected rcvChan to be closed, got open")
	}

	if !event3.WasCancelled() {
		t.Fatal("expected event3.WasCancelled to be true, got false")
	}

	if nh3.processed {
		t.Fatal("expected hn3.processed to be false, got true")
	}
}

func TestEnqueueTwice(t *testing.T) {
	queue := NewEventQueue("", 1)
	queue.Run()

	event := NewEvent(&FakeEvent{})
	res, err := queue.Enqueue(event)
	if err != nil {
		t.Fatalf("expected err == nil, got %v", err)
	}

	select {
	case <-res:
	case <-time.After(5 * time.Second):
		t.Fail()
	}

	res, err = queue.Enqueue(event)
	if res != nil {
		t.Fatalf("expected res to be nil, got %v", err)
	}

	if !errors.Is(err, ErrEventHasAlreadyEnqueueCalled) {
		t.Fatalf("expected err to be %v, got %v", ErrEventHasAlreadyEnqueueCalled, err)
	}

	queue.Stop()
	queue.WaitToBeDrained()
}

func TestForcefulDraining(t *testing.T) {
	queue := NewEventQueue("", 1)
	event := NewEvent(&FakeEvent{})
	res, err := queue.Enqueue(event)
	if err != nil {
		t.Fatalf("expected err to be nil, got %v", err)
	}

	if res == nil {
		t.Fatalf("expected res != nil, got nil")
	}

	queue.Stop()
	queue.WaitToBeDrained()

	select {
	case <-res:
	case <-time.After(3 * time.Second):
		t.Fail()
	}
}

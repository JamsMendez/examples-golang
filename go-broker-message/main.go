package main

import (
	"fmt"
	"sync"
	"time"
)

type Topic struct {
	name        string
	subscribers map[chan string]struct{}
	rwLock      sync.RWMutex
}

func NewTopic(name string) *Topic {
	return &Topic{
		name:        name,
		subscribers: map[chan string]struct{}{},
	}
}

func (t *Topic) Subscribe() chan string {
	ch := make(chan string)
	t.rwLock.Lock()
	t.subscribers[ch] = struct{}{}
	t.rwLock.Unlock()
	return ch
}

func (t *Topic) Unsubscribe(ch chan string) {
	t.rwLock.Lock()
	delete(t.subscribers, ch)
	t.rwLock.Unlock()
	close(ch)
}

func (t *Topic) Publish(msg string) {
	t.rwLock.RLock()
	defer t.rwLock.RUnlock()

	for subscribe := range t.subscribers {
		go func(ch chan string) {
			ch <- msg
		}(subscribe)
	}
}

func notifyMsg(name string, messages <-chan string) {
	for msg := range messages {
		fmt.Printf("%s: %s", name, msg)
	}
}

func main() {
	t := NewTopic("messages")

	jams := t.Subscribe()
	gopher := t.Subscribe()

	go notifyMsg("Jams", jams)
	go notifyMsg("Gopher", gopher)

	for i := 0; i < 5; i++ {
		ID := i + 1
		msg := fmt.Sprintf("Message %d\n", ID)
		t.Publish(msg)
		<-time.After(1 * time.Second)
	}
}

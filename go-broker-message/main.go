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
		subscribers: make(map[chan string]struct{}),
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

	for subscriber := range t.subscribers {
		go func(ch chan string) {
			ch <- msg
		}(subscriber)
	}
}

func readerMsg(name string, ch <-chan string, done chan struct{}) {
	for m := range ch {
		fmt.Printf("%s: %s\n", name, m)
	}
	fmt.Println(name, " unsubscribe")
	done <- struct{}{}
}

func main() {
	t := NewTopic("notify-golang")

	jams := t.Subscribe()
	gopher := t.Subscribe()

	done := make(chan struct{})

	go readerMsg("Jams", jams, done)
	go readerMsg("Gopher", gopher, done)

	for i := 0; i < 3; i++ {
		ID := i + 1
		msg := fmt.Sprintf("Message %d", ID)
		t.Publish(msg)
		<-time.After(1 * time.Second)
	}

	t.Unsubscribe(jams)
	t.Unsubscribe(gopher)

	for i := 0; i < 2; i++ {
		<-done
	}
	close(done)
}

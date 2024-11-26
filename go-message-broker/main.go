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

type Broker struct {
	topics map[string]*Topic
	rwLock sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string]*Topic),
	}
}

func (b *Broker) CreateTopic(name string) {
	b.rwLock.Lock()
	defer b.rwLock.Unlock()
	if _, exists := b.topics[name]; !exists {
		b.topics[name] = NewTopic(name)
	}
}

func (b *Broker) Topic(name string) (*Topic, bool) {
	b.rwLock.RLock()
	defer b.rwLock.RUnlock()
	t, exists := b.topics[name]
	return t, exists
}

func (b *Broker) PublishMsg(name, msg string) error {
	b.rwLock.RLock()
	defer b.rwLock.RUnlock()
	t, exists := b.topics[name]
	if !exists {
		return fmt.Errorf("topic %s not found", name)
	}

	t.Publish(msg)
	return nil
}

func main() {
	name := "messages"
	broker := NewBroker()
	broker.CreateTopic(name)

	t, ok := broker.Topic(name)
	if !ok {
		return
	}

	jams := t.Subscribe()
	gopher := t.Subscribe()

	go notifyMsg("Jams", jams)
	go notifyMsg("Gopher", gopher)

	for i := 0; i < 5; i++ {
		ID := i + 1
		msg := fmt.Sprintf("Message %d\n", ID)
		err := broker.PublishMsg(name, msg)
		if err != nil {
			fmt.Println(err)
		}
		<-time.After(1 * time.Second)
	}
}

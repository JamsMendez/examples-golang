package queue

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewQueue() *Queue {
	q := &Queue{}
	q.cond = sync.NewCond(&q.mu)

	return q
}

func (q *Queue) Enqueue(item int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.items = append(q.items, item)
	fmt.Printf("Produced: %d\n", item)

	q.cond.Signal()
}

func (q *Queue) Dequeue() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	fmt.Println("Pre for ...", len(q.items))
	for len(q.items) == 0 {
		fmt.Println("Wait ...")
		q.cond.Wait()
		fmt.Println("Post wait ...")
	}

	fmt.Println("Post for ...")
	<-time.After(2 * time.Second)

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

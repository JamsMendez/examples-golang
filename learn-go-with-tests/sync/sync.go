package sync

import (
	"sync"
)

type Counter struct {
	sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.Lock()
	defer c.Unlock()
	return c.value
}

/* type Counter struct {
	m     sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.m.Lock()
	defer c.m.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.value
} */

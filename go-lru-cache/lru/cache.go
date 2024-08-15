package lru

import (
	"container/list"
	"sync"
	"time"
)

type CacheItem struct {
	key       string
	value     any
	timestamp time.Time
}

type Cache struct {
	maxSize  int
	ttl      time.Duration
	cache    map[string]*list.Element
	eviction *list.List
	rwLock   sync.RWMutex
}

func NewCache(maxSize int, ttl time.Duration) *Cache {
	return &Cache{
		maxSize:  maxSize,
		ttl:      ttl,
		cache:    make(map[string]*list.Element),
		eviction: list.New(),
	}
}

func (c *Cache) Set(key string, value any) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.eviction.MoveToFront(ele)
		item := ele.Value.(*CacheItem)
		item.value = value
		item.timestamp = time.Now()
		return
	}

	if c.eviction.Len() >= c.maxSize {
		c.removeOldest()
	}

	item := &CacheItem{
		key:       key,
		value:     value,
		timestamp: time.Now(),
	}

	ele := c.eviction.PushFront(item)
	c.cache[key] = ele
}

func (c *Cache) Get(key string) (any, bool) {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()

	if ele, ok := c.cache[key]; ok {
		item := ele.Value.(*CacheItem)
		if time.Since(item.timestamp) > c.ttl {
			c.rwLock.RUnlock() // important unlock reading

			c.rwLock.Lock()
			c.eviction.Remove(ele)
			delete(c.cache, key)
			c.rwLock.Unlock()
			return nil, false
		}

		c.eviction.MoveToFront(ele)
		return item.value, true
	}

	return nil, false
}

func (c *Cache) Delete(key string) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.eviction.Remove(ele)
		delete(c.cache, key)
	}
}

func (c *Cache) removeOldest() {
	oldest := c.eviction.Back()
	if oldest != nil {
		c.eviction.Remove(oldest)
		item := oldest.Value.(*CacheItem)
		delete(c.cache, item.key)
	}
}

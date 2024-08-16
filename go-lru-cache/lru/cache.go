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
	// locker   sync.RWMutex
	mutex   sync.Locker
	rwMutex sync.Locker
}

func NewCache(maxSize int, ttl time.Duration, mutex, rwMutex sync.Locker) *Cache {
	return &Cache{
		maxSize:  maxSize,
		ttl:      ttl,
		cache:    make(map[string]*list.Element),
		eviction: list.New(),
		mutex:    mutex,
		rwMutex:  rwMutex,
	}
}

func (c *Cache) Set(key string, value any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

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

/* func (c *Cache) Get(key string) (any, bool) {
	c.rwLock.RLock()

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
		c.rwLock.RUnlock()
		return item.value, true
	}

	c.rwLock.RUnlock()
	return nil, false
} */

func (c *Cache) Get(key string) (any, bool) {
	c.rwMutex.Lock()

	if ele, ok := c.cache[key]; ok {
		item := ele.Value.(*CacheItem)
		if time.Since(item.timestamp) > c.ttl {
			c.rwMutex.Unlock() // important unlock reading

			c.mutex.Lock()
			c.eviction.Remove(ele)
			delete(c.cache, key)
			c.mutex.Unlock()
			return nil, false
		}

		c.eviction.MoveToFront(ele)
		c.rwMutex.Unlock()
		return item.value, true
	}

	c.rwMutex.Unlock()
	return nil, false
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

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

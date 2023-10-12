package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheMap sync.Map
}

type CacheItem struct {
	Value      any
	Expiration time.Time
}

func (c *Cache) Set(key string, value any, expiration time.Duration) {
	c.cacheMap.Store(key, CacheItem{Value: value, Expiration: time.Now().Add(expiration)})
}

func (c *Cache) Get(key string) (any, bool) {
	if val, ok := c.cacheMap.Load(key); ok {
		item := val.(CacheItem)
		if time.Now().Before(item.Expiration) {
			return item.Value, true
		}

		c.cacheMap.Delete(key)
	}

	return nil, false
}

func runCacheSimple() {
	cache := Cache{}

	cache.Set("JamsMendez", "something json", 5*time.Minute)

	if val, ok := cache.Get("JamsMendez"); ok {
		fmt.Println("Value found in cache", val)
	} else {
		fmt.Println("Value not found")
	}

	time.Sleep(6 * time.Minute)

	if val, ok := cache.Get("JamsMendez"); ok {
		fmt.Println("Value found in cache", val)
	} else {
		fmt.Println("Value not found in cache")
	}
}

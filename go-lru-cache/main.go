package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"

	"go-lru-cache/lru"
)

func main() {
	// LRU Cache
	// Least Recently Used
	// Cuando se necesite un nuevo elemento
	// se eliminada el menos usando recientemente.
	// Mutex vs RWMutex

	var rwMutex sync.RWMutex

	cache := lru.NewCache(10, 3*time.Second, &rwMutex, rwMutex.RLocker())

	tw := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")

	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))

		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, cache, &rwMutex, rwMutex.RLocker()), // RLocker return a rwMutex with RLock and RUnlock
			test(count, cache, &rwMutex, &rwMutex),
		)
	}
}

func producer(cache *lru.Cache, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 5; i > 0; i-- {
		cache.Set("users", fmt.Sprintf("user list %d", i))
		time.Sleep(1 * time.Nanosecond)
	}
}

func observer(cache *lru.Cache, wg *sync.WaitGroup) {
	defer wg.Done()
	_, _ = cache.Get("users")
}

func test(count int, cache *lru.Cache, mutex, rwMutex sync.Locker) time.Duration {
	var wg sync.WaitGroup

	nCount := count + 1

	wg.Add(nCount)

	beginTime := time.Now()

	go producer(cache, &wg) // nCount = 2 - 1 = 1

	for i := count; i > 0; i-- {
		go observer(cache, &wg)
	}

	wg.Wait()

	return time.Since(beginTime)
}

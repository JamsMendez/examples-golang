package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

/*
The sync package contains the concurrency primitives that are most useful
for low-level memory access synchronization. Critical section is the place
in your code that has access to a shared memory
*/

// Mutex
// Mutex stands for "mutual exclusion" and is a way to protect critical sections of your program
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	// bloquea el acceso
	c.mu.Lock()
	// libera el acceso al finalizar la function
	defer c.mu.Unlock()
	c.value++
}

// WaitGroup
// Call to add a group of goroutines
func runWaitGroup() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greeting", "good day"} {
		wg.Add(1)

		go func(s string) {
			defer wg.Done()
			fmt.Println(s)
		}(salutation)
	}

	wg.Wait()
}

// RWMutex
// More fine-grained memory control, being possible to request read-only lock
func runRWMutex(debug bool) {
	producer := func(key string, wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()

		for i := 5; i > 0; i-- {
			// logln(debug, "Producer.PreLock.i: ", key, i)
			l.Lock()
			logln(debug, "Producer: ", i, key)

			l.Unlock()
			// logln(debug, "Producer.PostUnlock.i: ", key, i)
			time.Sleep(1 * time.Nanosecond)
			// logln(debug, "Producer.PostUnlock.Finish.Iteration: ", key, i)
		}
	}

	observer := func(ID int, key string, wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		// logln(debug, "Observer.PreLock", ID, key)
		l.Lock()
		defer l.Unlock()

		logln(debug, "Observer: ", ID, key)
	}

	test := func(key string, count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup

		// count = 1
		nCount := count + 1
		// nCount = 2

		// logln(debug, "test.nCount: ", key, nCount)

		wg.Add(nCount)

		beginTestTime := time.Now()

		// logln(debug, "test.Pre.Producer", key)
		go producer(key, &wg, mutex) // nCount = 2 - 1 = 1
		// logln(debug, "test.Post.Producer", key)

		for i := count; i > 0; i-- {
			go observer(i, key, &wg, rwMutex)
		}

		// logln(debug, "test.PreWait: ", key, nCount)

		wg.Wait()

		// logln(debug, "test.PostWait: ", key, nCount)

		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")

	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))

		logln(debug, "MainRWMuter: ", count)

		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test("A", count, &m, m.RLocker()),
			test("B", count, &m, &m),
		)
	}
}

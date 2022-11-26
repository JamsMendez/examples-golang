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
func mainWaitGroup() {
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
func mainRWMutex() {
	producer := func(key string, wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()

		for i := 3; i > 0; i-- {
			fmt.Println("Producer.PreLock.i: ", key, i)
			l.Lock()
			fmt.Println("Producer.PostLock.i: ", key, i)
			l.Unlock()
			fmt.Println("Producer.PostUnlock.i: ", key, i)
			time.Sleep(1)
			fmt.Println("Producer.PostUnlock.Finish.Iteration: ", key, i)
		}
	}

	observer := func(key string, wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		fmt.Println("Observer.PreLock", key)
		l.Lock()
		fmt.Println("Observer.PostLock", key)
		defer l.Unlock()
	}

	test := func(key string, count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup

		nCount := count + 1

		fmt.Println("test.nCount: ", key, nCount)

		wg.Add(nCount)

		beginTestTime := time.Now()

		fmt.Println("test.Pre.Producer", key)
		go producer(key, &wg, mutex)
		fmt.Println("test.Post.Producer", key)

		for i := count; i > 0; i-- {
			fmt.Println("test.Pre.Observer: ", key, i)
			go observer(key, &wg, rwMutex)
			fmt.Println("test.Post.Observer: ", key, i)
		}

		fmt.Println("test.PreWait: ", key, nCount)

		wg.Wait()

		fmt.Println("test.PostWait: ", key, nCount)

		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")

	for i := 0; i < 3; i++ {
		count := int(math.Pow(2, float64(i)))

		fmt.Println("MainRWMuter: ", count)

		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test("A", count, &m, m.RLocker()),
			test("B", count, &m, &m),
		)

		fmt.Println("MainRWMuter.Finish.Iteration: ", count)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var a, b float64

func main() {
	fmt.Println("Mutex", "\t", "RWMutex")

	for i := 0; i < 10; i++ {
		done()
		time.Sleep(2 * time.Second)

		fmt.Println("Faster by: ", ((a-b)*100)/a, " %")
	}
}

func done() {
	r := rand.New(rand.NewSource(99))
	list1 := []int{}
	list2 := []int{}

	lock1 := sync.Mutex{}
	lock2 := sync.RWMutex{}

	wg1 := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}

	count := 1000

	for i := 0; i < 10; i++ {
		list1 = append(list1, i)
		list2 = append(list2, i)
	}

	// 1. spawn 1000 goroutines and synchronize through locks
	now1 := time.Now()
	for i := 0; i < count; i++ {
		wg1.Add(1)
		position := r.Intn(10)
		go func() {
			defer wg1.Done()
			lock1.Lock()

			defer lock1.Unlock()

			// goList, destroy automatically
			goList := []int{}
			// Process each element of list1 randomly
			for j := 0; j < count; j++ {
				goList = append(goList, list1[position])
			}
		}()
	}

	wg1.Wait()

	diff1 := time.Since(now1)

	// 2. spawn 1000 goroutines and synchronize through read-write locks
	now2 := time.Now()
	for i := 0; i < count; i++ {
		wg2.Add(1)
		position := r.Intn(10)
		go func() {
			defer wg2.Done()

			lock2.RLock()

			defer lock2.RUnlock()

			// goList, destroy automatically
			goList := []int{}
			// Process each element of list2 randomly
			for j := 0; j < count; j++ {
				goList = append(goList, list2[position])
			}
		}()
	}

	wg2.Wait()

	diff2 := time.Since(now2)
	fmt.Println(diff1, "\t", diff2)

	a = a + float64(diff1)
	b = b + float64(diff2)
}

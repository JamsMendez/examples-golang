package main

import (
	"fmt"
	"sync"
	"time"
)

// The Copy-on-Write pattern involves creating copies of shared data
// when modification is needed. This pattern ensures that each
// goroutine operates on its copy of the data, preventing
// interference between concurrent accesses.

type DataCopyAndWrite struct {
	sync.RWMutex
	Value string
}

func mainCopyAndWrite() {
	data := &DataCopyAndWrite{Value: "JamsMendez"}

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go mofifyDataCopyOnWrite(data, &wg)
	}

	wg.Wait()

	fmt.Println("All goroutines done!")
	fmt.Println("Final data value:", data.Value)
}

func mofifyDataCopyOnWrite(d *DataCopyAndWrite, wg *sync.WaitGroup) {
	defer wg.Done()

	// all goroutines will read the data value concurrently
	d.RLock()
	value := d.Value
	d.RUnlock()

	<-time.After(200 * time.Millisecond)

	// one goroutine will modify the data value at a time
	d.Lock()
	d.Value = value + " - Modified"
	d.Unlock()

	fmt.Println("Data modified by goroutine")
}

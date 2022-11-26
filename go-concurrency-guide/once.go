package main

import (
	"fmt"
	"sync"
)


func mainOnce() {
  var count int

  incrementCount := func ()  {
    count++
  }

  var once sync.Once

  var increments sync.WaitGroup
  increments.Add(100)

  for i := 0; i < 100; i++ {
    go func() {
      defer increments.Done()
      once.Do(incrementCount)
    }()
  }

  increments.Wait()

  fmt.Printf("Count is %d\n", count)
}

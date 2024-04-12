package main

import (
	"fmt"
	"sync"
	"time"
)

func mainPool() { 
  customPool := &sync.Pool{
    New: func() interface{} {
      now := time.Now().Format("2006-01-02 15:04:05")
      fmt.Println("Creating new instance.", now)
      time.Sleep(time.Second)

      // return struct{}{}
      return now
    },
  }

  // Get call new function defined in pool if there is no instance started
  customPool.Get()
  instance := customPool.Get()
  fmt.Println("Instance", instance)

  // here we put a previously retrieved instance back into the pool
  // this increases the number of instances available to one
  customPool.Put(instance)

  // when this call is executed, we will reuse the
  // previously allocated instance and put it back in the pool
  instance = customPool.Get()
  fmt.Println("InstanceAgain", instance)

  var numCalcsCreated int
  calcPool := &sync.Pool{
    New: func() interface{} {
      fmt.Println("New calc pool")

      numCalcsCreated += 1
      mem := make([]byte, 64)

      return &mem
    },
  }

  fmt.Println("calcPool.New", calcPool.New(), numCalcsCreated)

  calcPool.Put(calcPool.New())
  calcPool.Put(calcPool.New())
  calcPool.Put(calcPool.New())
  calcPool.Put(calcPool.New())

  fmt.Println("calcPool.Puts finished")

  calcPool.Get()

  const numWorkes = 64 * 64
  var wg sync.WaitGroup

  wg.Add(numWorkes)

  fmt.Println("calcPool.wg.Add ", numWorkes)

  for i := numWorkes; i > 0; i-- {
    go func (index int)  {
      defer wg.Done()

      mem := calcPool.Get().(*[]byte)
      defer calcPool.Put(mem)

      // Assume something interesting, but quick is being done with
      // this memory
    }(i)
  }

  wg.Wait()
  fmt.Printf("%d calculators were created.", numCalcsCreated)
}

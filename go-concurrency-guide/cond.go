package main

import (
	"fmt"
	"sync"
	"time"
)

/*
It would be better if there were some kind of way for a goroutine to efficiently sleep
until it was signaled to wake and check its condition. This is exactly what the Cond
type does for us.

The Cond and the Broadcast is the method that provides for notifying goroutines blocked
on Wait call that the condition has been triggered.
*/
type Button struct {
  Clicked *sync.Cond
}

func mainCond() {
  button := Button {
    Clicked: sync.NewCond(&sync.Mutex{}),
  }

  // running on goroutine every function that passed/registered
  // and wait, not exit until that that goroutine is confirmed
  // to be running
  subscribe := func(c *sync.Cond, param string, fn func(s string)) {
    var goroutineRunning sync.WaitGroup
    goroutineRunning.Add(1)

    go func(p string) {
      goroutineRunning.Done()
      c.L.Lock() // critical section
      defer c.L.Unlock()

      fmt.Println("Registered and wait ...") // LIFO
      c.Wait()
      fmt.Println("Running fn")
      fn(p)
    }(param)

    // fmt.Println(".................... GoroutineRunning.PRE.Wait")
    goroutineRunning.Wait()
    // fmt.Println(" .................... GoroutineRunning.POST.Wait")
  }

  var clickedRegistered sync.WaitGroup

  for _, v := range []string{
    "Maximizing window.",
    "Displaying annoying dialog box!",
    "Mouse clicked.",
  } {
    clickedRegistered.Add(1)

    subscribe(button.Clicked, v, func(s string) {
      fmt.Println(s)
      clickedRegistered.Done()
    })

    time.Sleep(time.Second)
  }

  fmt.Println("One moment ...")
  time.Sleep(time.Second)

  fmt.Println("Broadcast!")
  button.Clicked.Broadcast()

  fmt.Println("Wait clickedRegistered ...")
  clickedRegistered.Wait()
}


// A goroutine that is waiting for a signal, and a goroutine that is sending signals.
// Say we have a queue of fixed length 2, and 10 items we want to push onto the queue
func mainCoud() {
  c := sync.NewCond(&sync.Mutex{})

  queue := make([]interface{}, 0, 10)

  removeFromQueue := func (delay time.Duration)  {
    time.Sleep(delay)

    c.L.Lock()

    queue = queue[1:]

    fmt.Println("Removed from queue")

    c.L.Lock()
    c.Signal()
  }

  for i := 0; i < 10; i++ {
    c.L.Lock() // critical section

    // When the queue is equal to two the main goroutine is suspend
    // until a signal on the condition has been send
    for len(queue) == 2 {
      c.Wait()
    }

    fmt.Println("Adding to queue")

    queue = append(queue, struct{}{})

    go removeFromQueue(time.Second * 1)

    c.L.Unlock()
  }
}

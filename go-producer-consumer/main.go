package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
    for i := 0; i < 5; i++ {
       ch <- i + 1 
    }
    close(ch)
}


func consumer(ch chan int, done chan struct{}) {
    for v := range ch {
        fmt.Println("Consumidor: ", v) 
        <-time.After(1 * time.Second)
    }
    close(done)
}

func main() {
    queue := make(chan int)
    done := make(chan struct{})

    go producer(queue)
    go consumer(queue, done)

    <-done
}

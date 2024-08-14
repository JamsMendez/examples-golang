package main

import (
	"fmt"
	"go-lru-cache/lru"
	"time"
)

func main() {
    // LRU Cache
    // Least Recently Used
    // Cuando se necesite un nuevo elemento
    // se eliminada el menos usando recientemente. 
    // Mutex vs RWMutex

    cache := lru.NewCache(3, 3 * time.Second)
    cache.Set("users", "list users")
    fmt.Println(cache.Get("users"))

    <-time.After(4 * time.Second)

    fmt.Println(cache.Get("users"))


    cache.Set("users 1", "list users")
    <-time.After(time.Second)
    cache.Set("users 2", "list users")
    <-time.After(time.Second)
    cache.Set("users 3", "list users")
    <-time.After(time.Second)
    cache.Set("users 4", "list users")
    <-time.After(time.Second)

    fmt.Println(cache.Get("users 1"))
    fmt.Println(cache.Get("users 2"))
}

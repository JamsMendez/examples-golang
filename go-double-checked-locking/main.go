package main

import (
	"fmt"
	"sync"
)

var (
	instance *singleton
	mu       sync.Mutex

	once sync.Once
)

type singleton struct{}

func GetInstance() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			fmt.Println("Nueva instancia de singleton")
			instance = &singleton{}
		}
	}

	return instance
}

func GetOnceInstance() *singleton {
	once.Do(func() {
		fmt.Println("Nueva Instancia de singleton")
		instance = &singleton{}
	})

	if instance == nil {
		fmt.Println("Es nulo")
	}

	return instance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = GetOnceInstance()
		}()
	}

	wg.Wait()
}

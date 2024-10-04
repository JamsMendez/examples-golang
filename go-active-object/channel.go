package main

import (
	"fmt"
	"sync"
	"time"
)

// Struct con campo balance
type Account struct {
	balance int
}

func Run() {
	account := Account{balance: 100}
	balanceChan := make(chan int)

	// Goroutine para manejar la cuenta a través de un canal
	go func() {
		for newBalance := range balanceChan {
			account.balance += newBalance
		}
	}()

	// Simulamos múltiples operaciones de depósito concurrentes
	var wg sync.WaitGroup

	for i := 0; i < 30_000; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			var amount int
			if i%2 == 0 {
				amount = -100
			} else {
				amount = 100
			}

			deposit(balanceChan, amount)
		}(i)
	}

	wg.Wait()

	// 200
	fmt.Println("Balance: ", account.balance)

	time.Sleep(2 * time.Second) // Esperar a que las goroutines terminen

	close(balanceChan) // Cerrar el canal cuando ya no se necesite
}

func deposit(balanceChan chan int, amount int) {
	balanceChan <- amount
}

package main

import (
	"fmt"
	"sync"
)

type TicketReservation struct {
	availableSeats int
	mu             sync.Mutex
}

func (tr *TicketReservation) ReserveSeats(seats int) bool {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	if tr.availableSeats >= seats {
		fmt.Printf("Processing reservation for %d seats\n", seats)
		tr.availableSeats -= seats
		return true
	}

	return false
}

func (tr *TicketReservation) AvailableSeats() int {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	return tr.availableSeats
}

func runTicketReservation() {
	totalSeats := 21
	request := 10

	reservation := TicketReservation{availableSeats: totalSeats}

	var wg sync.WaitGroup

	for i := 0; i < request; i++ {
		wg.Add(1)

		go func(userID int) {
			defer wg.Done()

			seats := 2
			if reservation.ReserveSeats(seats) {
				fmt.Printf("User %d successfully reserved %d seats\n", userID, seats)
				return
			}

			fmt.Printf(
				"User %d failed to reverse seats. Available seats: %d",
				userID,
				reservation.AvailableSeats(),
			)
		}(i + 1)
	}

	wg.Wait()

	fmt.Println("Remainig available seats: ", reservation.AvailableSeats())
}

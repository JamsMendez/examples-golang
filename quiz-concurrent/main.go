package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	score := 0

	done := make(chan bool)

	go func() {
		for _, question := range getQuestions() {
			description, response := question[0], question[1]
			fmt.Printf("%s: ", description)
			_ = scanner.Scan()
			answer := strings.TrimSpace(scanner.Text())

			if answer == response {
				score++
			}
		}

		done <- true
	}()

	gameEngine(done, *time.NewTicker(10 * time.Second))

	fmt.Printf("Your score is %d!!!\n", score)
}

func gameEngine(done <-chan bool, ticker time.Ticker) {
	select {
	// Access to value <-done
	// case value := <-done:
	case <-done:
		fmt.Println("Finished ...")

	case <-ticker.C:
		fmt.Println("\nTimeout!!! ...")
	}
}

func getQuestions() (questions [][]string) {
	return [][]string{
		{"5 + 5", "10"},
		{"1 x 3", "3"},
		{"10 - 10", "0"},
		{"100 ÷ 10", "10"},
	}
}

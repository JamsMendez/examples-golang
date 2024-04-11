package main

import (
	"errors"
	"fmt"
)

type Result struct {
	Input any
	IsMod bool
	Err   error
}

func main() {
	inputStream := make(chan any)
	outputStream := make(chan Result)

	defer func() {
		close(inputStream)
		close(outputStream)
	}()

	inputs := []any{1, "Hola", 3, "JamsMendez", 5, 6, 7, 8}

	go producer(inputs, inputStream)
	go consumer(inputStream, outputStream)

	for i := 0; i < len(inputs); i++ {
		result := <-outputStream
		if result.Err != nil {
			fmt.Println(result.Err)
			continue
		}

		if result.IsMod {
			fmt.Println(result.Input, "is mod")
		} else {
			fmt.Println(result.Input, "is not mod")
		}
	}
}

func producer(inputs []any, inputStream chan<- any) {
	for _, input := range inputs {
		inputStream <- input
	}
}

func consumer(inputStream <-chan any, outputStream chan<- Result) {
	for input := range inputStream {
		result := Result{Input: input}
		ok, err := isMod(input)

		result.IsMod = ok
		result.Err = err

		outputStream <- result
	}
}

func isMod(input any) (bool, error) {
	switch t := input.(type) {
	case int:
		if t == 0 {
			return false, errors.New("zero is number not valid")
		}

		if t%2 == 0 {
			return true, nil
		}

		return false, fmt.Errorf("%d is not mod", t)

	case string:
		return false, errors.New("string is not valid")
	default:
		return false, errors.New("type is not valid")
	}
}

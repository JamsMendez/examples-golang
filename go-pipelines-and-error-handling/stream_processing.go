package main

import "fmt"

func runStreamProcessing() {
	inputs := []int{1, 2, 3, 4, 5}

	for _, input := range inputs {
		v := streamMultiply(streamAdd(input, 1),2)
		fmt.Println(v)
	}
}

func streamMultiply(input int, multiplier int) int {
	return input * multiplier
}

func streamAdd(input int, additive int) int {
	return input + additive
}

package main

import "fmt"

/*
The business logic has been kept simple to focus on the design pattern.

This kind of processing is called as batch processing, since the processing happens in batch. It has it’s own pros and cons.

Every stage created return another slice of equal length to input slice. This means at any point of a stage double the memory.

In this approach, the next stage starts only when all elements are done processing by the first stage.
*/

func runBatchProcessing() {
	input := []int{1, 2, 3, 4, 5}

	for _, r := range batchMultiply(batchAdd(input, 1), 2) {
		fmt.Println(r)
	}
}

func batchMultiply(numberStream []int, multiplier int) []int {
	result := make([]int, len(numberStream))

	for i, number := range numberStream {
		result[i] = number * multiplier
	}

	return result
}

func batchAdd(numberStream []int, additive int) []int {
	result := make([]int, len(numberStream))

	for i, number := range numberStream {
		result[i] = number + additive
	}

	return result
}

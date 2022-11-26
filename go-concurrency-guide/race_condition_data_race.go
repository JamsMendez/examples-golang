package main

import "fmt"

/*
Race condition occur when two or more operations must execute in the correct order,
but the program has not been written so that this order is guaranteed to maintained.

Data race is when one concurrent operation attempts to read a variable while at some
undetermined time another concurrent operation is attempting to write to the same
variable, the main func is the main goroutine.
*/

func mainRaceConditionDataRace() {
	var data int

	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("The value is %d\n", data)
	}
}

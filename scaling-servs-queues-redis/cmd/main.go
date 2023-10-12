package main

import (
	"fmt"
	"os"
	"scaling-servs-queues-redis/consumer"
	"scaling-servs-queues-redis/producer"
)

func main() {
	args := os.Args

	size := len(args)
	if size == 0 {
		fmt.Println("add argument: producer | consumer")
		return
	}

	value := args[size-1]

	switch value {
	case "consumer":
		consumer.Start()
	case "producer":
		producer.Start()
	default:
		fmt.Println("option invalid: use producer | consumer")
	}
}

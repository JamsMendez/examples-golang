package main

import (
	"context"
	"fmt"
	"time"
)

func TaskWithContextDeadline() {
	deadline := time.Now().Add(4 * time.Second)
	ctx, cancelCtx := context.WithDeadline(context.Background(), deadline)

	defer cancelCtx()

	printChannel := make(chan int)
	// background process
	go func(ctx context.Context, printChannel chan int) {
		for {
			select {
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					fmt.Println("background process ctx.Err()", err)
				}

				fmt.Println("finish background process")
				return
			case num := <-printChannel:
				fmt.Println("background process: ", num)
			}
		}
	}(ctx, printChannel)

	var stop bool
	for num := 0; num < 5; num++ {
		select {
		case printChannel <- num:
			time.Sleep(2 * time.Second)
		case <-ctx.Done():
			fmt.Println("main process done: ", num)
			stop = true
		}

		if stop {
			break
		}
	}

	cancelCtx()

	time.Sleep(time.Second)

	fmt.Println("finish main process")
}

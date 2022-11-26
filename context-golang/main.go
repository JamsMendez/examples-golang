package main

import (
	"context"
	"fmt"
	"time"
)

// Pasar el contexto en una funcion
func doSomething(ctx context.Context) {
	fmt.Println("Doing something")
}

// Se crea el contexto TODO o Background
// Background es para crear un contexto conocido
func createContext() {
	// ctx := context.TODO()
	ctx := context.Background()

	doSomething(ctx)
}

// ========================

const SOME_KEY = "some_key"

func doSomethingWithContextWithValue(ctx context.Context) {
	fmt.Printf("%s's value is %s\n", SOME_KEY, ctx.Value(SOME_KEY))
}

func createContextWithValue() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, SOME_KEY, "some_value...")

	doSomethingWithContextWithValue(ctx)
}

// =======================

func doSomethingContextAndAnother(ctx context.Context) {
	fmt.Printf("doSomethingContextAndAnother %s's value is %s\n", SOME_KEY, ctx.Value(SOME_KEY))

	anotherCtx := context.WithValue(ctx, SOME_KEY, "another value")
	doAnotherWithValue(anotherCtx)

	fmt.Printf("doSomethingContextAndAnother again %s's value is %s\n", SOME_KEY, ctx.Value(SOME_KEY))
}

func doAnotherWithValue(ctx context.Context) {
	fmt.Printf("doAnotherWithValue %s's value is %s\n", SOME_KEY, ctx.Value(SOME_KEY))
}

func createTwoContextWithValue() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, SOME_KEY, "some_value")

	doSomethingContextAndAnother(ctx)
}

// ==========================

func createContextWithContextDone() {
	ctx := context.Background()

	type WorkResult string
	resultsCh := make(chan *WorkResult)

	for {
		select {
		case <-ctx.Done():
			// The context is over, stop processing results
			return
		case result := <-resultsCh:
			fmt.Println("Result: ", result)
			// Process the results received
		}
	}
}

// ==========================

func doSomethingWithCancelAndAnother(ctx context.Context) {
	/*
		// Si context principal se cierra los demas se cierran
		// pero si es unos de los envueltos, solo aplica para si mismo
		ctx1, cancelCtx1 := context.WithCancel(ctx)
		ctx2, cancelCtx2 := context.WithCancel(ctx1)


		go func() {
			value := <-ctx1.Done()
			fmt.Println("Context One Done!!! ", value)
		}()


		go func() {
			value := <-ctx2.Done()
			fmt.Println("Context Two Done!!! ", value)
		}()

		time.Sleep(time.Second * 5)

		fmt.Println("CancelCtx2 ...")
		cancelCtx2()


		time.Sleep(time.Second * 5)

		fmt.Println("CancelCtx1 ...")
		cancelCtx1()

		time.Sleep(time.Second * 5)

		fmt.Println("Finish ...") */

	ctx, cancelCtx := context.WithCancel(ctx)
	printChannel := make(chan int)

	go doAnotherWithCancel(ctx, printChannel)

	for num := 1; num <= 3; num++ {
		printChannel <- num
	}

	cancelCtx()

	time.Sleep(time.Second * 1)

	fmt.Println("doSomethingWithCancelAndAnother: finishe ...")
}

func doAnotherWithCancel(ctx context.Context, printChannel chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Println("doAnotherWithCancel.ERROR: ", err)
			}
			fmt.Println("doAnotherWithCancel finished ...")
			return
		case num := <-printChannel:
			fmt.Println("doAnotherWithCancel ", num)
		}
	}
}

func createTwoContextWithCancel() {
	ctx := context.Background()
	doSomethingWithCancelAndAnother(ctx)
}

// ==========================

// Context con una fecha de limite de vida, de igual forma se puede cancelar de manera manual
func doSomethingContextWithDeadline(ctx context.Context) {
	deadline := time.Now().Add(time.Second * 2)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	// Aqui solo especificas el tiempo limite de vida del contexto
	// ctx, cancelCtx := context.WithTimeout(ctx, time.Second * 2)
	defer cancelCtx()

	printChannel := make(chan int)
	go doAnotherWithDeadline(ctx, printChannel)

	for num := 1; num <= 3; num++ {
		select {
		case printChannel <- num:
			time.Sleep(time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	time.Sleep(time.Second)

	fmt.Println("doSomethingContextWithDeadline finished ...")
}

func doAnotherWithDeadline(ctx context.Context, printChannel chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Println("doAnotherWithDeadline.ERROR: ", err)
			}
			fmt.Println("doAnotherWithDeadline finished ...")
			return
		case num := <-printChannel:
			fmt.Println("doAnotherWithDeadline ", num)
		}
	}
}

func createContextWithDeadline() {
	ctx := context.Background()
	doSomethingContextWithDeadline(ctx)
}

// ==========================

func main() {
	// createContext()
	// createContextWithValue()
	// createTwoContextWithValue()
	// createTwoContextWithCancel()
	createContextWithDeadline()
}

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type ContextKey string

const keyTraceID ContextKey = "trace-id"

// main and backgroun process share same context
func TaskOneContext(wg *sync.WaitGroup) {
	ctx := context.WithValue(context.Background(), keyTraceID, "req-117")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		// fmt.Println("background.Done: ", <-ctx.Done())
		time.Sleep(2 * time.Second)

		fmt.Println("background process ctx.Err(): ", ctx.Err())
		fmt.Printf("background process ctx.Value(\"%s\"): %s\n", keyTraceID, ctx.Value(keyTraceID))
		fmt.Println("finish background process ...")
	}(ctx)

	fmt.Println("Finish main process")
}

// main and backgroun process has distinc context
func TaskTwoContext(wg *sync.WaitGroup) {
	ctx := context.WithValue(context.Background(), keyTraceID, "req-058")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg.Add(1)
	go func(ctx context.Context) {
		bCtx := context.WithValue(context.Background(), keyTraceID, ctx.Value(keyTraceID))
		defer wg.Done()

		// fmt.Println("background.Done: ", <-ctx.Done())
		time.Sleep(2 * time.Second)

		fmt.Println("background process bCtx.Err()", bCtx.Err())
		fmt.Printf("background process bCtx.Value(\"%s\"): %s\n", keyTraceID, bCtx.Value(keyTraceID))
		fmt.Println("finish background process ...")
	}(ctx)

	fmt.Println("Finish main process")
}

// Create custom context
type contextWithoutDeadline struct {
	ctx context.Context
}

func (c *contextWithoutDeadline) Deadline() (time.Time, bool) {
	return c.ctx.Deadline()
}

func (c *contextWithoutDeadline) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *contextWithoutDeadline) Err() error {
	return fmt.Errorf("ContextCustom.ERROR: %v", c.ctx.Err())
}

func (c *contextWithoutDeadline) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func TaskWithCustonContext(wg *sync.WaitGroup) {
	ctx := context.WithValue(context.Background(), keyTraceID, "req-089")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg.Add(1)
	go func(ctx context.Context) {
		bCtx := &contextWithoutDeadline{ctx}
		defer wg.Done()

		time.Sleep(time.Second)

		fmt.Println("background process bCtx.Err(): ", bCtx.Err())
		fmt.Printf("background process bCtx.Value(\"%s\"): %s\n", keyTraceID, bCtx.Value(keyTraceID))
		fmt.Println("finish background process")
	}(ctx)

	fmt.Println("finish main process")
}

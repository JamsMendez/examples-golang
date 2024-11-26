package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	done := make(chan struct{})
	logCh := make(chan string)
	errCh := make(chan error)
	dataCh := make(chan LogData)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go readFile(ctx, "log.txt", logCh, errCh)
	go processLog(ctx, logCh, dataCh, errCh)
	go storeLog(ctx, dataCh, errCh, done)

	select {
	case err := <-errCh:
		log.Println("pipeline error: ", err)
	case <-time.After(5 * time.Second):
		log.Println("pipeline timeout")
	case <-done:
		log.Println("pipeline success")
	}
}

type LogData struct {
	Original  string
	Processed string
	Timestamp time.Time
}

func readFile(ctx context.Context, file string, out chan<- string, errCh chan<- error) {
	defer close(out)

	f, err := os.Open(file)
	if err != nil {
		errCh <- fmt.Errorf("open file error: %w", err)
		return
	}

	defer func() {
		err = f.Close()
		if err != nil {
			log.Println("close file error: ", err)
		}
	}()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			errCh <- fmt.Errorf("reading canceled: %w", ctx.Err())
			return
		case out <- scanner.Text():
			<-time.After(500 * time.Millisecond)
		}
	}

	if err = scanner.Err(); err != nil {
		errCh <- fmt.Errorf("scanner file error: %w", err)
	}
}

func processLog(ctx context.Context, in <-chan string, out chan<- LogData, errCh chan<- error) {
	defer close(out)

	for line := range in {
		select {
		case <-ctx.Done():
			errCh <- fmt.Errorf("processing canceled: %w", ctx.Err())
			return
		default:
			nLine := fmt.Sprintf("[INFO] %s", line)

			out <- LogData{
				Original:  line,
				Processed: nLine,
				Timestamp: time.Now(),
			}
		}
	}
}

func storeLog(ctx context.Context, in <-chan LogData, errCh chan<- error, done chan struct{}) {
	defer close(done)

	for data := range in {
		select {
		case <-ctx.Done():
			errCh <- fmt.Errorf("storing canceled: %w", ctx.Err())
			return
		default:
			fmt.Printf("%s: %s\n", data.Processed, data.Timestamp.Format(time.DateTime))
		}
	}
}

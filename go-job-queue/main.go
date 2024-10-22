package main

import (
	"fmt"
	"time"

	"go-job-queue/jobqueue"
)

type ExampleJob struct {
	Name string
}

func (ej *ExampleJob) Execute() {
	fmt.Println("Executing Job: ", ej.Name)
	<-time.After(1 * time.Second)
}

func main() {
	queue := jobqueue.New(10)

	queue.Enqueue(&ExampleJob{Name: "Job 1"})
	queue.Enqueue(&ExampleJob{Name: "Job 2"})
	queue.Enqueue(&ExampleJob{Name: "Job 3"})

	queue.Close()
}

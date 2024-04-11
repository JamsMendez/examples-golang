package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID      string
	Product string
	Price   int
}

type OrderError struct {
	Order Order
	Err   error
}

type Pipeline struct {
	Orders chan Order
	Errors chan OrderError
}

func runPipelineDataChannel() {
	pipeline := NewPipeline()

	go pipeline.StartProcessing()
	go pipeline.StartErrorListening()

	orders := []Order{
		{"1", "Shoes", 25},
		{"5", "Books", 1000},
		{"1", "TV", 200},
		{"1", "Sit", 75},
	}

	for _, order := range orders {
		pipeline.Orders <- order
	}

	time.Sleep(10 * time.Second)

	close(pipeline.Orders)
	close(pipeline.Errors)
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		Orders: make(chan Order, 10),
		Errors: make(chan OrderError, 10),
	}
}

func (p *Pipeline) ProcessOrder(o Order) {
	time.Sleep(time.Second)

	if o.Price%2 == 0 {
		p.Errors <- OrderError{
			Order: o,
			Err:   fmt.Errorf("cannot process even price of %s", o.Product),
		}
		return
	}

	fmt.Printf("Orders %s processed successfully\n", o.ID)
}

func (p *Pipeline) StartProcessing() {
	for order := range p.Orders {
		go p.ProcessOrder(order)
	}
}

func (p *Pipeline) StartErrorListening() {
	for err := range p.Errors {
		fmt.Printf("Error processing order %s: %v\n", err.Order.ID, err.Err)
	}
}

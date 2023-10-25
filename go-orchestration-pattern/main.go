package main

import "fmt"

type OrderService struct{}

type InventaryService struct{}

type ShippginService struct{}

type OrderOrchestrator struct {
	orderService     *OrderService
	inventaryService *InventaryService
	shippginService  *ShippginService
}

func (os *OrderService) PlaceOrder() {
	fmt.Println("Order placed.")
}

func (is *InventaryService) ProcessOrder() {
	fmt.Println("Order processed in Inventory.")
}

func (ss *ShippginService) ShipOrder() {
	fmt.Println("Order shipped.")
}

func (oo *OrderOrchestrator) FulFillOrder() {
	oo.orderService.PlaceOrder()
	oo.inventaryService.ProcessOrder()
	oo.shippginService.ShipOrder()
}

func main() {
	orderService := &OrderService{}
	inventaryService := &InventaryService{}
	shippingService := &ShippginService{}

	orderOrchestrator := &OrderOrchestrator{
		orderService:     orderService,
		inventaryService: inventaryService,
		shippginService:  shippingService,
	}

	orderOrchestrator.FulFillOrder()
}

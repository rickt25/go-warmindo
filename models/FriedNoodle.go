package models

import (
	"fmt"
	"time"
)

type FriedNoodle struct {
	Flavor     string
	Price      uint64
	Stock      uint16
	Toppings   []NoodleTopping
	SpiceLevel int
}

func (n FriedNoodle) GetName() string {
	return "Indomie Goreng " + n.Flavor
}

func (n FriedNoodle) GetPrice() uint64 {
	return n.Price
}

func (n FriedNoodle) GetStock() uint16 {
	return n.Stock
}

func (n FriedNoodle) Cook() {
	fmt.Printf("Cooking %s...\n", n.GetName())
	time.Sleep(2 * time.Second)
	fmt.Println("Cooking done.", n.GetName())
}

func (n FriedNoodle) PrintDetail() {
	fmt.Printf("Name: %s\n", n.GetName())
	fmt.Printf("Price: %d\n", n.GetPrice())
	fmt.Printf("Stock: %d\n", n.GetStock())
}

func (n FriedNoodle) AddStock(amount int) {
	n.Stock += uint16(amount)
}

func (n FriedNoodle) ReduceStock(amount int) {
	if n.Stock < uint16(amount) {
		fmt.Println("Not enough stock")
		return
	}
	n.Stock -= uint16(amount)
}

func (n FriedNoodle) AddTopping(topping NoodleTopping) {
	n.Toppings = append(n.Toppings, topping)
}

func (n FriedNoodle) GetToppings() []NoodleTopping {
	return n.Toppings
}

package models

import (
	"fmt"
	"time"
)

type BoiledNoodle struct {
	Flavor     string
	Price      uint64
	Stock      uint16
	Toppings   []NoodleTopping
	SpiceLevel int
}

func (n BoiledNoodle) GetName() string {
	return "Indomie Rebus " + n.Flavor
}

func (n BoiledNoodle) GetPrice() uint64 {
	return n.Price
}

func (n BoiledNoodle) GetStock() uint16 {
	return n.Stock
}

func (n BoiledNoodle) Cook() {
	fmt.Printf("Cooking %s...\n", n.GetName())
	time.Sleep(2 * time.Second)
	fmt.Println("Cooking done.", n.GetName())
}

func (n BoiledNoodle) PrintDetail() {
	fmt.Printf("Name: %s\n", n.GetName())
	fmt.Printf("Price: %d\n", n.GetPrice())
	fmt.Printf("Stock: %d\n", n.GetStock())
}

func (n BoiledNoodle) AddStock(amount int) {
	n.Stock += uint16(amount)
}

func (n BoiledNoodle) ReduceStock(amount int) {
	if n.Stock < uint16(amount) {
		fmt.Println("Not enough stock")
	}
	n.Stock -= uint16(amount)
}

func (n BoiledNoodle) AddTopping(topping NoodleTopping) {
	n.Toppings = append(n.Toppings, topping)
}

func (n BoiledNoodle) GetToppings() []NoodleTopping {
	return n.Toppings
}

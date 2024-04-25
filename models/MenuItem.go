package models

type MenuItem interface {
	GetName() string
	GetPrice() uint64
	GetStock() uint16
	PrintDetail()
	AddStock(amount int)
	ReduceStock(amount int)
	Cook()
}

type Toppingable interface {
	MenuItem
	AddTopping(topping NoodleTopping)
	GetToppings() []NoodleTopping
}

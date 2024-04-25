package models

type Order struct {
	Menu       MenuItem
	Toppings   []NoodleTopping
	SpiceLevel int
}

func (order Order) CalculateTotal() uint64 {
	total := order.Menu.GetPrice()

	for _, topping := range order.Toppings {
		total += topping.Price
	}

	return total
}

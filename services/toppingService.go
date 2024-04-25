package services

import (
	"fmt"
	"warmindoProject/models"
)

type ToppingService struct {
	toppings []models.NoodleTopping
}

func NewToppingService() *ToppingService {
	return &ToppingService{toppings: []models.NoodleTopping{}}
}

func SeedTopping(toppingService *ToppingService) {
	toppingList := []models.NoodleTopping{
		{Name: "Telur", Price: 4000},
		{Name: "Kornet", Price: 6000},
		{Name: "Bakso", Price: 5000},
		{Name: "Sosis", Price: 5000},
	}

	for _, topping := range toppingList {
		toppingService.AddTopping(topping)
	}
}

func (toppingService *ToppingService) AddTopping(topping models.NoodleTopping) {
	toppingService.toppings = append(toppingService.toppings, topping)
}

func (toppingService ToppingService) PrintToppings() {
	fmt.Println("Noodle Toppings:")
	for i, topping := range toppingService.toppings {
		fmt.Printf("%d. %s (%d)\n", i+1, topping.Name, topping.Price)
	}
}

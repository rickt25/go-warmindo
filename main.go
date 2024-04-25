package main

import (
	"fmt"
	"warmindoProject/services"
)

type Topping struct {
	name  string
	price uint64
}

var EXIT_MENU_NUMBER int = 0

func main() {
	// Initialize menus
	menuService := services.NewMenuService()
	services.SeedMenu(menuService)

	// Initialize toppings
	toppingService := services.NewToppingService()
	services.SeedTopping(toppingService)

	// Initialize order service
	orderService := services.NewOrderService()

	var menu int
	for {
		printMenu()
		fmt.Print(">> ")
		fmt.Scan(&menu)

		if menu == EXIT_MENU_NUMBER {
			break
		}

		switch menu {
		case 1:
			menuService.PrintMenuByType()
		case 2:
			toppingService.PrintToppings()
		case 3:
			orderService.PlaceOrder(menuService, toppingService)
		case 4:
			orderService.PrintOrderList()
		}
	}

	fmt.Println("Program exited")
}

func printMenu() {
	fmt.Println("=======================")
	fmt.Println("|     Warmindo 76     |")
	fmt.Println("=======================")
	fmt.Println("| 1. Indomie List     |")
	fmt.Println("| 2. Topping List     |")
	fmt.Println("| 3. Order Indomie    |")
	fmt.Println("| 4. Transaction List |")
	fmt.Println("| 0. Exit             |")
	fmt.Println("=======================")
}

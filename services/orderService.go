package services

import (
	"fmt"
	"warmindoProject/models"
)

type OrderService struct {
	orders []models.Order
}

func NewOrderService() *OrderService {
	return &OrderService{orders: []models.Order{}}
}

func (orderService *OrderService) PlaceOrder(menuService *MenuService, toppingService *ToppingService) {
	order := orderService.createOrder(menuService, toppingService)
	orderService.addToOrderList(order)
	orderService.printReceipt(order)
}

func (orderService *OrderService) PrintOrderList() {
	fmt.Println("================================================")
	fmt.Println(" TRANSACTION LIST")
	fmt.Println("================================================")
	for i, order := range orderService.orders {
		fmt.Printf("Order %d:\n", i+1)
		fmt.Println("Noodle:", order.Menu.GetName())
		fmt.Println("Spice Level:", order.SpiceLevel)
		fmt.Println("Toppings:")
		for _, topping := range order.Toppings {
			fmt.Println("\t -", topping.Name)
		}
		fmt.Println("Order Total:", order.CalculateTotal())
		fmt.Println("================================================")
	}
}

func (orderService *OrderService) createOrder(menuService *MenuService, toppingService *ToppingService) models.Order {
	var flavor string
	var toppings []models.NoodleTopping
	spiceLevel := -1

	menuService.PrintMenuByType()
	flavor = orderService.chooseNoodleFlavor(menuService)

	spiceLevel = orderService.chooseSpiceLevel()

	toppings = orderService.chooseToppings(toppingService)

	selectedNoodle := orderService.getSelectedNoodle(menuService, flavor)
	order := models.Order{
		Menu:       selectedNoodle,
		Toppings:   toppings,
		SpiceLevel: spiceLevel,
	}

	return order
}

func (orderService *OrderService) chooseNoodleFlavor(menuService *MenuService) string {
	var flavor string
	isNoodleExist := false
	for !isNoodleExist {
		fmt.Print("Choose Indomie Flavor: ")
		fmt.Scan(&flavor)
		for _, item := range menuService.menuItems {
			switch noodle := item.(type) {
			case models.BoiledNoodle:
				if noodle.Flavor == flavor {
					isNoodleExist = true
				}
			case models.FriedNoodle:
				if noodle.Flavor == flavor {
					isNoodleExist = true
				}
			}
		}
		if !isNoodleExist {
			fmt.Println("Invalid flavor. Please choose again.")
		}
	}
	return flavor
}

func (orderService *OrderService) chooseSpiceLevel() int {
	spiceLevel := -1
	for spiceLevel < 0 || spiceLevel > 5 {
		fmt.Print("Choose Spice Level (0 to 5): ")
		fmt.Scan(&spiceLevel)
		if spiceLevel < 0 || spiceLevel > 5 {
			fmt.Println("Spice Level only from 0 to 5")
		}
	}
	return spiceLevel
}

func (orderService *OrderService) chooseToppings(toppingService *ToppingService) []models.NoodleTopping {
	var toppings []models.NoodleTopping
	toppingService.PrintToppings()
	fmt.Println("Choose Toppings (enter 0 to finish):")
	for {
		var toppingIndex int
		fmt.Print("Topping Index: ")
		fmt.Scan(&toppingIndex)
		if toppingIndex == 0 {
			break
		}
		if toppingIndex >= 1 && toppingIndex <= len(toppingService.toppings) {
			fmt.Println(toppingService.toppings[toppingIndex-1].Name, "added.")
			toppings = append(toppings, toppingService.toppings[toppingIndex-1])
		} else {
			fmt.Println("Invalid topping index")
		}
	}
	return toppings
}

func (orderService *OrderService) getSelectedNoodle(menuService *MenuService, flavor string) models.MenuItem {
	var selectedNoodle models.MenuItem
	for i, item := range menuService.menuItems {
		switch noodle := item.(type) {
		case models.BoiledNoodle:
			if noodle.Flavor == flavor {
				selectedNoodle = noodle
				noodle.Stock--
				menuService.menuItems[i] = noodle
			}
		case models.FriedNoodle:
			if noodle.Flavor == flavor {
				selectedNoodle = noodle
				noodle.Stock--
				menuService.menuItems[i] = noodle
			}
		}
	}
	return selectedNoodle
}

func (orderService *OrderService) addToOrderList(order models.Order) {
	orderService.orders = append(orderService.orders, order)
}

func (orderService *OrderService) printReceipt(order models.Order) {
	fmt.Println("================================================")
	fmt.Println(" RECEIPT")
	fmt.Println("================================================")
	fmt.Println("Noodle:", order.Menu.GetName())
	fmt.Println("Spice Level:", order.SpiceLevel)
	fmt.Println("Toppings:")
	for _, topping := range order.Toppings {
		fmt.Println("\t -", topping.Name)
	}
	fmt.Println("================================================")
	fmt.Println(" Order Total:", order.CalculateTotal())
	fmt.Println("================================================")
}

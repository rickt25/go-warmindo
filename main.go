package main

import "fmt"

type Topping struct {
	name  string
	price uint64
}

type Noodle interface {
	cook()
}

type BoiledNoodle struct {
	flavor string
	price  uint64
	stock  uint16
}

type FriedNoodle struct {
	flavor string
	price  uint64
	stock  uint16
}

type Order struct {
	noodle     Noodle
	toppings   []Topping
	spiceLevel int
}

func (order Order) calculateTotal() uint64 {
	var price uint64

	// Check if the noodle is a BoiledNoodle
	switch order.noodle.(type) {
	case BoiledNoodle:
		boiledNoodle := order.noodle.(BoiledNoodle)
		price = boiledNoodle.price
	case FriedNoodle:
		friedNoodle := order.noodle.(FriedNoodle)
		price = friedNoodle.price
	default:
		fmt.Println("Noodle not found")
	}

	// Add the price of toppings
	for _, topping := range order.toppings {
		price += topping.price
	}

	return price
}

func (noodle BoiledNoodle) cook() {
	fmt.Println("Boiling Noodle...")
}

func (noodle FriedNoodle) cook() {
	fmt.Println("Frying Noodle...")
}

var noodleList []Noodle
var toppingList []Topping
var orderList []Order

func initData() {
	noodleList = append(noodleList, BoiledNoodle{"Soto", 11000, 10})
	noodleList = append(noodleList, BoiledNoodle{"Kari Ayam", 10000, 10})
	noodleList = append(noodleList, BoiledNoodle{"Ayam Bawang", 10000, 10})
	noodleList = append(noodleList, FriedNoodle{"Original", 10000, 10})
	noodleList = append(noodleList, FriedNoodle{"Jumbo", 16000, 10})
	noodleList = append(noodleList, FriedNoodle{"Rendang", 11000, 10})

	toppingList = append(toppingList, Topping{"Telur", 4000})
	toppingList = append(toppingList, Topping{"Kornet", 6000})
	toppingList = append(toppingList, Topping{"Bakso", 5000})
	toppingList = append(toppingList, Topping{"Sosis", 5000})
}

func main() {
	initData()
	var menu int

	for {
		fmt.Println("Warmindo 76")
		fmt.Println("1. Indomie List")
		fmt.Println("2. Add Indomie")
		fmt.Println("3. Topping List")
		fmt.Println("4. Add Topping")
		fmt.Println("5. Order Indomie")
		fmt.Println("6. Transaction List")
		fmt.Println("0. Exit")
		fmt.Print(">> ")
		fmt.Scan(&menu)

		if menu == 0 {
			break
		}

		switch menu {
		case 1:
			printIndomieList()
		case 2:
			addIndomie()
		case 3:
			printToppingList()
		case 4:
			addTopping()
		case 5:
			orderIndomie()
		case 6:
			printOrderList()
		}
	}

	fmt.Println("Program exited")
}

func printToppingList() {
	fmt.Println("======================")
	for index, t := range toppingList {
		fmt.Printf("%d.", index+1)
		fmt.Printf("\tName: %s\n", t.name)
		fmt.Printf("\tPrice: %d\n", t.price)
		fmt.Println("======================")
	}
}

func addTopping() {
	var name string
	var price uint64

	for name == "" {
		fmt.Print("Topping Name: ")
		fmt.Scan(&name)
	}

	for price < 1 || price > 10000 {
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price < 1 {
			fmt.Println("Price can't be less than 0")
		}

		if price > 10000 {
			fmt.Println("Price can't be greater than 10000")
		}
	}

	toppingList = append(toppingList, Topping{name, price})
	fmt.Println("New Topping added")
}

func addIndomie() {
	var flavor string
	var price, stock, noodleType int

	for noodleType < 1 || noodleType > 2 {
		fmt.Println("Pilih Jenis Indomie [1|2]: ")
		fmt.Println("1. Goreng")
		fmt.Println("2. Rebus")
		fmt.Print("> ")
		fmt.Scan(&noodleType)
	}

	for flavor == "" {
		fmt.Print("Rasa Indomie: ")
		fmt.Scan(&flavor)
	}

	for price < 1 || price > 10000 {
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price < 1 {
			fmt.Println("Price can't be less than 0")
		}

		if price > 10000 {
			fmt.Println("Price can't be greater than 10000")
		}
	}

	for stock < 1 || stock > 1000 {
		fmt.Print("Stock: ")
		fmt.Scan(&stock)

		if price < 1 {
			fmt.Println("Stock can't be less than 0")
		}

		if price > 1000 {
			fmt.Println("Stock	 can't be greater than 1000")
		}
	}

	switch noodleType {
	case 1:
		noodleList = append(noodleList, FriedNoodle{flavor, uint64(price), uint16(stock)})
	case 2:
		noodleList = append(noodleList, BoiledNoodle{flavor, uint64(price), uint16(stock)})
	}

	fmt.Println("New Indomie added")
	return
}

func printIndomieList() {
	fmt.Println("===========================================")
	for i, n := range noodleList {
		if boiledNoodle, ok := n.(BoiledNoodle); ok {
			fmt.Printf("%d.", i+1)
			fmt.Printf("\tName: Indomie Rebus %s\n", boiledNoodle.flavor)
			fmt.Printf("\tPrice: %d\n", boiledNoodle.price)
			fmt.Printf("\tStock: %d\n", boiledNoodle.stock)
			fmt.Println("===========================================")
		} else if friedNoodle, ok := n.(FriedNoodle); ok {
			fmt.Printf("%d.", i+1)
			fmt.Printf("\tName: Indomie Goreng %s\n", friedNoodle.flavor)
			fmt.Printf("\tPrice: %d\n", friedNoodle.price)
			fmt.Printf("\tStock: %d\n", friedNoodle.stock)
			fmt.Println("===========================================")
		}

	}
}

func orderIndomie() {
	var flavor string
	var toppings []Topping
	spiceLevel := -1
	isNoodleExist := false

	printIndomieList()

	for !isNoodleExist {
		fmt.Print("Choose Indomie Flavor: ")
		fmt.Scan(&flavor)

		// Check if the specified flavor exists
		var foundNoodle Noodle
		for _, noodle := range noodleList {
			switch n := noodle.(type) {
			case BoiledNoodle:
				if n.flavor == flavor {
					isNoodleExist = true
					foundNoodle = n
					break
				}
			case FriedNoodle:
				if n.flavor == flavor {
					isNoodleExist = true
					foundNoodle = n
					break
				}
			}
		}

		if foundNoodle == nil {
			fmt.Println("Indomie with the specified flavor not found")
		}
	}

	for spiceLevel < 0 || spiceLevel > 5 {
		fmt.Print("Choose Spice Level (0 to 5): ")
		fmt.Scan(&spiceLevel)

		if spiceLevel < 0 || spiceLevel > 5 {
			fmt.Println("Spice Level only from 0 to 5")
		}
	}

	printToppingList()

	fmt.Println("Choose Toppings (enter 0 to finish):")
	for {
		var toppingIndex int
		fmt.Print("Topping Index: ")
		fmt.Scan(&toppingIndex)
		if toppingIndex == 0 {
			break
		}
		if toppingIndex >= 1 && toppingIndex <= len(toppingList) {
			fmt.Println(toppingList[toppingIndex-1].name, "added.")
			toppings = append(toppings, toppingList[toppingIndex-1])
		} else {
			fmt.Println("Invalid topping index")
		}
	}

	var selectedNoodle Noodle
	for _, noodle := range noodleList {
		switch n := noodle.(type) {
		case BoiledNoodle:
			if n.flavor == flavor {
				selectedNoodle = noodle
				break
			}
		case FriedNoodle:
			if n.flavor == flavor {
				selectedNoodle = noodle
				break
			}
		}
	}

	order := Order{
		noodle:     selectedNoodle,
		toppings:   toppings,
		spiceLevel: spiceLevel,
	}

	orderList = append(orderList, order)

	fmt.Println("================================================")
	fmt.Println(" RECEIPT")
	fmt.Println("================================================")
	for _, noodle := range noodleList {
		switch n := noodle.(type) {
		case BoiledNoodle:
			if n.flavor == flavor {
				fmt.Println(" Indomie Rebus", n.flavor)
				fmt.Println(" SpiceLevel:", order.spiceLevel)
				break
			}
		case FriedNoodle:
			if n.flavor == flavor {
				fmt.Println(" Indomie Goreng", n.flavor)
				fmt.Println(" SpiceLevel:", order.spiceLevel)
				break
			}
		}
	}
	fmt.Println(" Toppings:")
	for _, topping := range toppings {
		fmt.Println("\t -", topping.name)
	}
	fmt.Println("================================================")
	fmt.Println(" Order Total:", order.calculateTotal())
	fmt.Println("================================================")
}

func printOrderList() {
	if len(orderList) == 0 {
		fmt.Println("No orders yet")
		return
	}

	fmt.Println("================================================")
	fmt.Println(" TRANSACTION LIST")
	fmt.Println("================================================")
	for i, order := range orderList {
		fmt.Println(" Order", i+1)
		fmt.Println(" Noodle:")
		for _, noodle := range noodleList {
			if noodle == order.noodle {
				switch n := noodle.(type) {
				case BoiledNoodle:
					fmt.Println("\tIndomie Rebus", n.flavor)
				case FriedNoodle:
					fmt.Println("\tIndomie Goreng", n.flavor)
				}
			}
		}
		fmt.Println(" Toppings:")
		for _, topping := range order.toppings {
			fmt.Printf("\t- %s\n", topping.name)
		}
		fmt.Println(" Spice Level:", order.spiceLevel)
		fmt.Println(" Order Total:", order.calculateTotal())
		fmt.Println("================================================")
	}
}

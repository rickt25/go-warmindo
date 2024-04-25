package main

import "fmt"

type Topping struct {
	name  string
	price uint64
}

type Data interface {
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
	noodle     Data
	toppings   []Topping
	spiceLevel int
}

func (order Order) calculateTotal() uint64 {
	var p uint64

	switch order.noodle.(type) {
	case BoiledNoodle:
		boiledNoodle := order.noodle.(BoiledNoodle)
		p = boiledNoodle.price
	case FriedNoodle:
		friedNoodle := order.noodle.(FriedNoodle)
		p = friedNoodle.price
	default:
		fmt.Println("Noodle not found")
	}

	for _, topping := range order.toppings {
		p += topping.price
	}

	return p
}

func (n BoiledNoodle) cook() {
	fmt.Println("Boiling Noodle...")
}

func (n FriedNoodle) cook() {
	fmt.Println("Frying Noodle...")
}

var nList []Data    // List of Noodle
var tList []Topping // List of Topping
var tran []Order    // List of Transaction

func first() {
	nList = append(nList, BoiledNoodle{"Soto", 11000, 10})
	nList = append(nList, BoiledNoodle{"Kari Ayam", 10000, 10})
	nList = append(nList, BoiledNoodle{"Ayam Bawang", 10000, 10})
	nList = append(nList, FriedNoodle{"Original", 10000, 10})
	nList = append(nList, FriedNoodle{"Jumbo", 16000, 10})
	nList = append(nList, FriedNoodle{"Rendang", 11000, 10})

	tList = append(tList, Topping{"Telur", 4000})
	tList = append(tList, Topping{"Kornet", 6000})
	tList = append(tList, Topping{"Bakso", 5000})
	tList = append(tList, Topping{"Sosis", 5000})
}

func main() {
	first()
	var menu int

	for {
		fmt.Println("Warmindo 76")
		fmt.Println("1. Indomie List")
		fmt.Println("2. Topping List")
		fmt.Println("3. Order Indomie")
		fmt.Println("4. Transaction List")
		fmt.Println("0. Exit")
		fmt.Print(">> ")
		fmt.Scan(&menu)

		if menu == 0 {
			break
		}

		switch menu {
		case 1:
			fmt.Println("===========================================")
			for i, n := range nList {
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
		case 2:
			fmt.Println("======================")
			for index, t := range tList {
				fmt.Printf("%d.", index+1)
				fmt.Printf("\tName: %s\n", t.name)
				fmt.Printf("\tPrice: %d\n", t.price)
				fmt.Println("======================")
			}
		case 3:
			var flavor string
			var toppings []Topping
			spiceLevel := -1
			isItThere := false // Check if Noodle Exist

			fmt.Println("===========================================")
			for i, noodle := range nList {
				switch n := noodle.(type) {
				case BoiledNoodle:
					fmt.Printf("%d.", i+1)
					fmt.Printf("\tName: Indomie Rebus %s\n", n.flavor)
					fmt.Printf("\tPrice: %d\n", n.price)
					fmt.Printf("\tStock: %d\n", n.stock)
					fmt.Println("===========================================")
				case FriedNoodle:
					fmt.Printf("%d.", i+1)
					fmt.Printf("\tName: Indomie Goreng %s\n", n.flavor)
					fmt.Printf("\tPrice: %d\n", n.price)
					fmt.Printf("\tStock: %d\n", n.stock)
					fmt.Println("===========================================")
				}
			}

			for !isItThere {
				fmt.Print("Choose Indomie Flavor: ")
				fmt.Scan(&flavor)

				// Check if noodle found
				var foundNoodle Data
				for _, noodle := range nList {
					switch n := noodle.(type) {
					case BoiledNoodle:
						if n.flavor == flavor {
							if n.stock > 0 {
								isItThere = true
								foundNoodle = n
							} else {
								fmt.Println("Indomie out of stock")
							}

							break
						}
					case FriedNoodle:
						if n.flavor == flavor {
							if n.stock > 0 {
								isItThere = true
								foundNoodle = n
							} else {
								fmt.Println("Indomie out of stock")
							}

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

			fmt.Println("======================")
			for i, t := range tList {
				fmt.Printf("%d.", i+1)
				fmt.Printf("\tName: %s\n", t.name)
				fmt.Printf("\tPrice: %d\n", t.price)
				fmt.Println("======================")
			}

			fmt.Println("Choose Toppings (enter 0 to finish):")
			for {
				var toppingIndex int
				fmt.Print("Topping Index: ")
				fmt.Scan(&toppingIndex)
				if toppingIndex == 0 {
					break
				}
				if toppingIndex >= 1 && toppingIndex <= len(tList) {
					fmt.Println(tList[toppingIndex-1].name, "added.")
					toppings = append(toppings, tList[toppingIndex-1])
				} else {
					fmt.Println("Invalid topping index")
				}
			}

			var selectedNoodle Data
			for i, noodle := range nList {
				switch n := noodle.(type) {
				case BoiledNoodle:
					if n.flavor == flavor {
						selectedNoodle = noodle
						n.stock--
						nList[i] = n
						break
					}
				case FriedNoodle:
					if n.flavor == flavor {
						selectedNoodle = noodle
						n.stock--
						nList[i] = n
						break
					}
				}
			}

			order := Order{
				noodle:     selectedNoodle,
				toppings:   toppings,
				spiceLevel: spiceLevel,
			}

			tran = append(tran, order)

			fmt.Println("================================================")
			fmt.Println(" RECEIPT")
			fmt.Println("================================================")
			for _, noodle := range nList {
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
		case 4:
			if len(tran) == 0 {
				fmt.Println("No orders yet")
				return
			}

			fmt.Println("================================================")
			fmt.Println(" TRANSACTION LIST")
			fmt.Println("================================================")
			for i, order := range tran {
				fmt.Println(" Order", i+1)
				fmt.Println(" Noodle:")
				for _, noodle := range nList {
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
	}

	fmt.Println("Program exited")
}

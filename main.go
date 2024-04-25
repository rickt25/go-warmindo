package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"time"
)

type Topping struct {
	name  string
	price uint64
}

type Cook func()

type Noodle struct {
	noodleType string
	flavor     string
	price      uint64
	stock      uint16
	cook       Cook
}

type Order struct {
	noodle     Noodle
	toppings   []Topping
	spiceLevel int
}

var nList []Noodle  // List of Noodle
var tList []Topping // List of Topping
var tran []Order    // List of Transaction

func first() {
	nList = append(nList, Noodle{"Rebus", "Soto", 11000, 10, func() {
		fmt.Println("Boiling the noodle...")
		// simulate cooking process
		time.Sleep(2 * time.Second)
		fmt.Println("Finish boiling")
	}})
	nList = append(nList, Noodle{"Rebus", "Kari Ayam", 10000, 10, func() {
		fmt.Println("Boiling the noodle...")
		// simulate cooking process
		time.Sleep(2 * time.Second)
		fmt.Println("Finish boiling")
	}})
	nList = append(nList, Noodle{"Rebus", "Ayam Bawang", 10000, 10, func() {
		fmt.Println("Boiling the noodle...")
		// simulate cooking process
		time.Sleep(2 * time.Second)
		fmt.Println("Finish boiling")
	}})
	nList = append(nList, Noodle{"Goreng", "Original", 10000, 10, func() {
		fmt.Println("Boiling the noodle...")
		// simulate cooking process
		time.Sleep(2 * time.Second)
		fmt.Println("Finish boiling")
	}})
	nList = append(nList, Noodle{"Goreng", "Jumbo", 16000, 10, func() {
		fmt.Println("Boiling the noodle...")
		// simulate cooking process
		time.Sleep(2 * time.Second)
		fmt.Println("Finish boiling")
	}})
	nList = append(nList, Noodle{"Goreng", "Rendang", 11000, 10, func() {
		fmt.Println("Boiling the noodle...")
		// simulate cooking process
		time.Sleep(2 * time.Second)
		fmt.Println("Finish boiling")
	}})

	tList = append(tList, Topping{"Telur", 4000})
	tList = append(tList, Topping{"Kornet", 6000})
	tList = append(tList, Topping{"Bakso", 5000})
	tList = append(tList, Topping{"Sosis", 5000})
}

func main() {
	first()
	var menu int

	for {
		fmt.Println("=======================")
		fmt.Println("|     Warmindo 76     |")
		fmt.Println("=======================")
		fmt.Println("| 1. Indomie List     |")
		fmt.Println("| 2. Topping List     |")
		fmt.Println("| 3. Order Indomie    |")
		fmt.Println("| 4. Transaction List |")
		fmt.Println("| 0. Exit             |")
		fmt.Println("=======================")
		fmt.Print(">> ")
		fmt.Scan(&menu)

		if menu == 0 {
			break
		}

		switch menu {
		case 1:
			fmt.Println("===========================================")
			for i, n := range nList {
				if n.noodleType == "Goreng" {
					fmt.Printf("%d.", i+1)
					fmt.Printf("\tName: Indomie Rebus %s\n", n.flavor)
					fmt.Printf("\tPrice: %d\n", n.price)
					fmt.Printf("\tStock: %d\n", n.stock)
					fmt.Println("===========================================")
				} else if n.noodleType == "Rebus" {
					fmt.Printf("%d.", i+1)
					fmt.Printf("\tName: Indomie Goreng %s\n", n.flavor)
					fmt.Printf("\tPrice: %d\n", n.price)
					fmt.Printf("\tStock: %d\n", n.stock)
					fmt.Println("===========================================")
				}
			}
		case 2:
			fmt.Println("=======================")
			for index, t := range tList {
				fmt.Printf("%d.", index+1)
				fmt.Printf("\tName: %s\n", t.name)
				fmt.Printf("\tPrice: %d\n", t.price)
				fmt.Println("=======================")
			}
		case 3:
			test()
		case 4:
			if len(tran) == 0 {
				fmt.Println("No orders yet")
			}

			fmt.Println("================================================")
			fmt.Println(" TRANSACTION LIST")
			fmt.Println("================================================")
			for i, order := range tran {
				fmt.Println(" Order", i+1)
				fmt.Println(" Noodle:")
				for _, noodle := range nList {
					if reflect.DeepEqual(noodle, order.noodle) {
						switch noodle.noodleType {
						case "Rebus":
							fmt.Println("\tIndomie Rebus", noodle.flavor)
						case "Goreng":
							fmt.Println("\tIndomie Goreng", noodle.flavor)
						}
					}
				}
				fmt.Println(" Toppings:")
				for _, topping := range order.toppings {
					fmt.Printf("\t- %s\n", topping.name)
				}
				fmt.Println(" Spice Level:", order.spiceLevel)
				//count total of orderan
				total := order.noodle.price
				for _, topping := range order.toppings {
					total += topping.price
				}
				fmt.Println(" Order Total:", total)
				fmt.Println("================================================")
			}
		}
	}

	fmt.Println("Program exited")
}

func test() {
	var flavor string
	var toppings []Topping
	spiceLevel := -1
	isItThere := false // Check if Noodle Exist

	fmt.Println("===========================================")
	for i, noodle := range nList {
		switch noodle.noodleType {
		case "Rebus":
			fmt.Printf("%d.", i+1)
			fmt.Printf("\tName: Indomie Rebus %s\n", noodle.flavor)
			fmt.Printf("\tPrice: %d\n", noodle.price)
			fmt.Printf("\tStock: %d\n", noodle.stock)
			fmt.Println("===========================================")
		case "Goreng":
			fmt.Printf("%d.", i+1)
			fmt.Printf("\tName: Indomie Goreng %s\n", noodle.flavor)
			fmt.Printf("\tPrice: %d\n", noodle.price)
			fmt.Printf("\tStock: %d\n", noodle.stock)
			fmt.Println("===========================================")
		}
	}

	for !isItThere {
		fmt.Print("Choose Indomie Flavor: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		flavor = scanner.Text()

		// Check if noodle found
		foundNoodle := Noodle{}
		for _, noodle := range nList {
			switch noodle.noodleType {
			case "Rebus":
				if noodle.flavor == flavor {
					if noodle.stock > 0 {
						isItThere = true
						foundNoodle = noodle
					} else {
						fmt.Println("Indomie out of stock")
					}

					break
				}
			case "Goreng":
				if noodle.flavor == flavor {
					if noodle.stock > 0 {
						isItThere = true
						foundNoodle = noodle
					} else {
						fmt.Println("Indomie out of stock")
					}

					break
				}
			default:
				fmt.Println("Indomie variant not found")
			}
		}

		if reflect.DeepEqual(foundNoodle, nList) {
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

	fmt.Println("=======================")
	for i, t := range tList {
		fmt.Printf("%d.", i+1)
		fmt.Printf("\tName: %s\n", t.name)
		fmt.Printf("\tPrice: %d\n", t.price)
		fmt.Println("=======================")
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

	// selected noodle?
	var sNoodle Noodle
	for i, noodle := range nList {
		if noodle.flavor == flavor {
			sNoodle = noodle
			noodle.stock--
			nList[i] = noodle
			break
		}
	}

	// cooking mama 👨‍🍳🍜🍜🍜🍜
	sNoodle.cook()

	order := Order{
		noodle:     sNoodle,
		toppings:   toppings,
		spiceLevel: spiceLevel,
	}

	tran = append(tran, order)

	fmt.Println("================================================")
	fmt.Println(" RECEIPT")
	fmt.Println("================================================")
	for _, noodle := range nList {
		switch noodle.noodleType {
		case "Rebus":
			if noodle.flavor == flavor {
				fmt.Println(" Indomie Rebus", noodle.flavor)
				fmt.Println(" SpiceLevel:", order.spiceLevel)
				break
			}
		case "Goreng":
			if noodle.flavor == flavor {
				fmt.Println(" Indomie Goreng", noodle.flavor)
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
	//count total of orderan
	total := sNoodle.price
	for _, topping := range toppings {
		total += topping.price
	}
	fmt.Println(" Order Total:", total)
	fmt.Println("================================================")
}

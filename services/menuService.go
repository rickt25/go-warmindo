package services

import (
	"fmt"
	"reflect"
	"warmindoProject/models"
)

type MenuService struct {
	menuItems []models.MenuItem
}

func NewMenuService() *MenuService {
	return &MenuService{menuItems: []models.MenuItem{}}
}

func SeedMenu(menuService *MenuService) {
	menuList := []models.MenuItem{
		models.BoiledNoodle{Flavor: "Soto", Price: 11000, Stock: 10},
		models.BoiledNoodle{Flavor: "Kari Ayam", Price: 10000, Stock: 10},
		models.BoiledNoodle{Flavor: "Ayam Bawang", Price: 10000, Stock: 10},
		models.FriedNoodle{Flavor: "Original", Price: 9000, Stock: 10},
		models.FriedNoodle{Flavor: "Jumbo", Price: 15000, Stock: 10},
		models.FriedNoodle{Flavor: "Rendang", Price: 11000, Stock: 10},
	}

	for _, item := range menuList {
		menuService.AddMenuItem(item)
	}
}

func (menuService *MenuService) AddMenuItem(item models.MenuItem) {
	menuService.menuItems = append(menuService.menuItems, item)
}

func (menuService MenuService) PrintMenuByType() {
	menuByType := make(map[string][]models.MenuItem)

	for _, item := range menuService.menuItems {
		itemType := reflect.TypeOf(item).Name()
		menuByType[itemType] = append(menuByType[itemType], item)
	}

	for menuType, items := range menuByType {
		fmt.Println("================================================")
		fmt.Println(menuType + ":")
		fmt.Println("================================================")
		for _, item := range items {
			item.PrintDetail()
			fmt.Printf("=================================\n")
		}
	}
}

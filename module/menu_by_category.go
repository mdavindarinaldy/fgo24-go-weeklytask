package module

import (
	"sync"
)

type MenuByCategory struct {
	foodItems    []Item
	drinkItems   []Item
	snackItems   []Item
	dessertItems []Item
}

func NewMenuByCategory() *MenuByCategory {
	return &MenuByCategory{
		foodItems:    []Item{},
		drinkItems:   []Item{},
		snackItems:   []Item{},
		dessertItems: []Item{},
	}
}

func (mc *MenuByCategory) FilterFood(wg *sync.WaitGroup, menu *Menu) {
	defer wg.Done()
	mc.foodItems = []Item{}
	for _, item := range menu.GetAll() {
		if item.Category == "food" {
			mc.foodItems = append(mc.foodItems, item)
		}
	}
}

func (mc *MenuByCategory) FilterDrink(wg *sync.WaitGroup, menu *Menu) {
	defer wg.Done()
	mc.drinkItems = []Item{}
	for _, item := range menu.GetAll() {
		if item.Category == "drink" {
			mc.drinkItems = append(mc.drinkItems, item)
		}
	}
}

func (mc *MenuByCategory) FilterSnack(wg *sync.WaitGroup, menu *Menu) {
	defer wg.Done()
	mc.snackItems = []Item{}
	for _, item := range menu.GetAll() {
		if item.Category == "snack" {
			mc.snackItems = append(mc.snackItems, item)
		}
	}
}

func (mc *MenuByCategory) FilterDessert(wg *sync.WaitGroup, menu *Menu) {
	defer wg.Done()
	mc.dessertItems = []Item{}
	for _, item := range menu.GetAll() {
		if item.Category == "dessert" {
			mc.dessertItems = append(mc.dessertItems, item)
		}
	}
}

func (mc *MenuByCategory) EmptyItems(wg *sync.WaitGroup) {
	defer wg.Done()
	mc.foodItems = []Item{}
	mc.drinkItems = []Item{}
	mc.snackItems = []Item{}
	mc.dessertItems = []Item{}
}

func (mc *MenuByCategory) GetFoodItems() []Item {
	return mc.foodItems
}

func (mc *MenuByCategory) GetDrinkItems() []Item {
	return mc.drinkItems
}

func (mc *MenuByCategory) GetSnackItems() []Item {
	return mc.snackItems
}

func (mc *MenuByCategory) GetDessertItems() []Item {
	return mc.dessertItems
}

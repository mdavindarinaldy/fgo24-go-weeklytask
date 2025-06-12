package module

import (
	"sync"
)

var FoodItems []Item = []Item{}

var DrinkItems []Item = []Item{}

var SnackItems []Item = []Item{}

var DessertItems []Item = []Item{}

func FoodItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.Category == "Food" {
			FoodItems = append(FoodItems, item)
		}
	}
}

func DrinkItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.Category == "Drink" {
			DrinkItems = append(DrinkItems, item)
		}
	}
}

func SnackItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.Category == "Snack" {
			SnackItems = append(SnackItems, item)
		}
	}
}

func DessertItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.Category == "Dessert" {
			DessertItems = append(DessertItems, item)
		}
	}
}

func EmptyItemByCategory(wg *sync.WaitGroup) {
	defer wg.Done()
	FoodItems = []Item{}
	DrinkItems = []Item{}
	SnackItems = []Item{}
	DessertItems = []Item{}
}

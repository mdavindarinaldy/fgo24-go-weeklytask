package module

import "sync"

var NusantaraItems []Item = []Item{}
var WesternItems []Item = []Item{}
var JapaneseItems []Item = []Item{}

func NusantaraItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.Origin == "Nusantara" {
			NusantaraItems = append(NusantaraItems, item)
		}
	}
}

func WesternItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.Origin == "Western" {
			WesternItems = append(WesternItems, item)
		}
	}
}

func JapaneseItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.Origin == "Japanese" {
			JapaneseItems = append(JapaneseItems, item)
		}
	}
}

func EmptyItemByOrigin(wg *sync.WaitGroup) {
	defer wg.Done()
	FoodItems = []Item{}
	DrinkItems = []Item{}
	SnackItems = []Item{}
	DessertItems = []Item{}
}

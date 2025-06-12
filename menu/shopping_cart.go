package menu

import "sync"

type Item struct {
	name     string
	price    int
	origin   string
	category string
}

var Items []Item = []Item{
	{
		name:     "Nasi Goreng",
		price:    15000,
		origin:   "Nusantara",
		category: "Food",
	}, {
		name:     "Nasi Uduk",
		price:    10000,
		origin:   "Nusantara",
		category: "Food",
	}, {
		name:     "Fish and Chips",
		price:    18000,
		origin:   "Western",
		category: "Food",
	}, {
		name:     "Sushi",
		price:    20000,
		origin:   "Japanese",
		category: "Food",
	}, {
		name:     "Es Dawet",
		price:    9000,
		origin:   "Nusantara",
		category: "Drink",
	}, {
		name:     "Cherry Blossom Tea",
		price:    5000,
		origin:   "Japanese",
		category: "Drink",
	}, {
		name:     "Kombucha",
		price:    8000,
		origin:   "Japanese",
		category: "Drink",
	}, {
		name:     "Jasmine Tea in Affogato Latte",
		price:    13000,
		origin:   "Western",
		category: "Drink",
	}, {
		name:     "Mochi",
		price:    9000,
		origin:   "Japanese",
		category: "Snack",
	}, {
		name:     "Dango",
		price:    10000,
		origin:   "Japanese",
		category: "Snack",
	}, {
		name:     "Pisang Goreng",
		price:    5000,
		origin:   "Nusantara",
		category: "Snack",
	}, {
		name:     "Serabi",
		price:    7500,
		origin:   "Nusantara",
		category: "Snack",
	}, {
		name:     "Cheesecake",
		price:    20000,
		origin:   "Western",
		category: "Dessert",
	}, {
		name:     "Red Velvet Cake",
		price:    22000,
		origin:   "Western",
		category: "Dessert",
	}, {
		name:     "Creme Brulee",
		price:    18000,
		origin:   "Western",
		category: "Dessert",
	}, {
		name:     "Bubur Kacang Ijo",
		price:    10000,
		origin:   "Nusantara",
		category: "Dessert",
	},
}

var FoodItems []Item = []Item{}

var DrinkItems []Item = []Item{}

var SnackItems []Item = []Item{}

var DessertItems []Item = []Item{}

var Cart []Item

func AddItem(item Item) {
	Cart = append(Cart, item)
}

func FoodItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.category == "Food" {
			FoodItems = append(FoodItems, item)
		}
	}
}

func DrinkItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.category == "Food" {
			DrinkItems = append(DrinkItems, item)
		}
	}
}

func SnackItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.category == "Food" {
			SnackItems = append(SnackItems, item)
		}
	}
}

func DessertItem(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, item := range Items {
		if item.category == "Food" {
			DessertItems = append(DessertItems, item)
		}
	}
}

// implementasi goroutines : lakukan pencarian berdasarkan range harga, rating, kategori --> masukkan ke dalam

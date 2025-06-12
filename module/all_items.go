package module

type Item struct {
	Name     string
	Price    int
	Origin   string
	Category string
}

var Items []Item = []Item{
	{
		Name:     "Nasi Goreng",
		Price:    15000,
		Origin:   "Nusantara",
		Category: "Food",
	}, {
		Name:     "Nasi Uduk",
		Price:    10000,
		Origin:   "Nusantara",
		Category: "Food",
	}, {
		Name:     "Fish and Chips",
		Price:    18000,
		Origin:   "Western",
		Category: "Food",
	}, {
		Name:     "Sushi",
		Price:    20000,
		Origin:   "Japanese",
		Category: "Food",
	}, {
		Name:     "Es Dawet",
		Price:    9000,
		Origin:   "Nusantara",
		Category: "Drink",
	}, {
		Name:     "Cherry Blossom Tea",
		Price:    5000,
		Origin:   "Japanese",
		Category: "Drink",
	}, {
		Name:     "Kombucha",
		Price:    8000,
		Origin:   "Japanese",
		Category: "Drink",
	}, {
		Name:     "Jasmine Tea in Affogato Latte",
		Price:    13000,
		Origin:   "Western",
		Category: "Drink",
	}, {
		Name:     "Mochi",
		Price:    9000,
		Origin:   "Japanese",
		Category: "Snack",
	}, {
		Name:     "Dango",
		Price:    10000,
		Origin:   "Japanese",
		Category: "Snack",
	}, {
		Name:     "Pisang Goreng",
		Price:    5000,
		Origin:   "Nusantara",
		Category: "Snack",
	}, {
		Name:     "Serabi",
		Price:    7500,
		Origin:   "Nusantara",
		Category: "Snack",
	}, {
		Name:     "Cheesecake",
		Price:    20000,
		Origin:   "Western",
		Category: "Dessert",
	}, {
		Name:     "Red Velvet Cake",
		Price:    22000,
		Origin:   "Western",
		Category: "Dessert",
	}, {
		Name:     "Creme Brulee",
		Price:    18000,
		Origin:   "Western",
		Category: "Dessert",
	}, {
		Name:     "Bubur Kacang Ijo",
		Price:    10000,
		Origin:   "Nusantara",
		Category: "Dessert",
	},
}

func AddItem(item Item) {
	Items = append(Items, item)
}

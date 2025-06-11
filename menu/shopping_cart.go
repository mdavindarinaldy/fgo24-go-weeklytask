package menu

type Item struct {
	name   string
	price  int
	origin string
}

var FoodItems []Item = []Item{
	{
		name:   "Nasi Goreng",
		price:  15000,
		origin: "Nusantara",
	}, {
		name:   "Nasi Uduk",
		price:  10000,
		origin: "Nusantara",
	}, {
		name:   "Fish and Chips",
		price:  18000,
		origin: "Western",
	}, {
		name:   "Sushi",
		price:  20000,
		origin: "Japanese",
	},
}

var DrinkItems []Item = []Item{
	{
		name:   "Es Dawet",
		price:  9000,
		origin: "Nusantara",
	}, {
		name:   "Cherry Blossom Tea",
		price:  5000,
		origin: "Japanese",
	}, {
		name:   "Kombucha",
		price:  8000,
		origin: "Japanese",
	}, {
		name:   "Jasmine Tea in Affogato Latte",
		price:  13000,
		origin: "Western",
	},
}

var SnackItems []Item = []Item{
	{
		name:   "Mochi",
		price:  9000,
		origin: "Japanese",
	}, {
		name:   "Dango",
		price:  10000,
		origin: "Japanese",
	}, {
		name:   "Pisang Goreng",
		price:  5000,
		origin: "Nusantara",
	}, {
		name:   "Serabi",
		price:  7500,
		origin: "Nusantara",
	},
}

var DessertItems []Item = []Item{
	{
		name:   "Cheesecake",
		price:  20000,
		origin: "Western",
	}, {
		name:   "Red Velvet Cake",
		price:  22000,
		origin: "Western",
	}, {
		name:   "Creme Brulee",
		price:  18000,
		origin: "Western",
	}, {
		name:   "Bubur Kacang Ijo",
		price:  10000,
		origin: "Nusantara",
	},
}

var Cart []Item

func AddItem(item Item) {
	Cart = append(Cart, item)
}

// implementasi goroutines : lakukan pencarian berdasarkan range harga, rating, kategori --> masukkan ke dalam

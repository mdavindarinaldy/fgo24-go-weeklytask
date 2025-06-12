package module

var Cart []Item = []Item{}

func AddCartItem(item Item) {
	Cart = append(Cart, item)
}

// implementasi goroutines : lakukan pencarian berdasarkan range harga, rating, kategori --> masukkan ke dalam

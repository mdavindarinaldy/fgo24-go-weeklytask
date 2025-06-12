package module

import (
	"fmt"
	"time"
)

var Cart []Item = []Item{}

func AddCartItem(item Item) {
	Cart = append(Cart, item)
	fmt.Printf("\n%s berhasil ditambahkan ke dalam keranjang!\n\n", item.Name)
	time.Sleep(time.Second)
}

// implementasi goroutines : lakukan pencarian berdasarkan range harga, rating, kategori --> masukkan ke dalam

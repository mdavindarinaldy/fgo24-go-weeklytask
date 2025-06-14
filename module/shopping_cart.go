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

func RemoveCartItem(index int) {
	if index == 1 {
		Cart = Cart[1:]
	} else if index == len(Cart) {
		Cart = Cart[:index-1]
	} else {
		Cart = append(Cart[:index-1], Cart[index:]...)
	}
}

func EmptyCartItem() {
	Cart = []Item{}
}

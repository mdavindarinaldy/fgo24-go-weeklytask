package menu

import (
	"fmt"
	"time"
)

func ShoppingCart() {
	fmt.Println("Keranjang :")
	var totalPrice int = 0
	for i, item := range Cart {
		fmt.Printf("%d. %s - Rp.%d\n", i+1, item.name, item.price)
		totalPrice = totalPrice + item.price
	}
	fmt.Printf("Total Pembayaran : Rp.%d\n", totalPrice)
	time.Sleep(time.Duration(3) * time.Second)
}

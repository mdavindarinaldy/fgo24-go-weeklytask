package menu

import (
	"fgo24-weekly-go/module"
	"fmt"
	"time"
)

func ShoppingCart() {
	fmt.Println("Keranjang :")
	var totalPrice int = 0
	for i, item := range module.Cart {
		fmt.Printf("%d. %s - Rp.%d\n", i+1, item.Name, item.Price)
		totalPrice = totalPrice + item.Price
	}
	fmt.Printf("Total Pembayaran : Rp.%d\n", totalPrice)
	time.Sleep(time.Duration(3) * time.Second)
}

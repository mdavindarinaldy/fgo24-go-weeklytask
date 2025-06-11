package menu

import (
	"fmt"
	"os"
)

func CheckOut() {
	fmt.Println("Checkout :")
	var totalPrice int = 0
	for i, item := range Cart {
		fmt.Printf("%d. %s - Rp.%d\n", i+1, item.name, item.price)
		totalPrice = totalPrice + item.price
	}
	fmt.Printf("Total Pembayaran : Rp.%d\n", totalPrice)
	fmt.Println("Terima kasih sudah berbelanja")
	defer os.Exit(0)
}

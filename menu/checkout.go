package menu

import (
	"fgo24-weekly-go/module"
	"fmt"
	"os"
)

func CheckOut() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Checkout :")
	var totalPrice int = 0
	for i, item := range module.Cart {
		fmt.Printf("%d. %s - Rp.%d\n", i+1, item.Name, item.Price)
		totalPrice = totalPrice + item.Price
	}
	fmt.Printf("Total Pembayaran : Rp.%d\n", totalPrice)
	fmt.Println("Terima kasih sudah berbelanja")
	defer os.Exit(0)
}

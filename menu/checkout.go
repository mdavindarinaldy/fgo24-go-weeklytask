package menu

import (
	"fgo24-weekly-go/module"
	"fgo24-weekly-go/utils"
	"fmt"
	"time"
)

func CheckOut() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Checkout :")

	if len(module.Cart) == 0 {
		fmt.Println("Belum ada item yang dimasukkan ke dalam keranjang!")
		fmt.Printf("\n[DIKEMBALIKAN KE MENU UTAMA]\n\n")
		time.Sleep(2 * time.Second)
		return
	}

	var totalPrice int = 0
	for i, item := range module.Cart {
		fmt.Printf("%d. %s - Rp.%d\n", i+1, item.Name, item.Price)
		totalPrice = totalPrice + item.Price
	}
	fmt.Printf("Total Pembayaran : Rp.%d\n", totalPrice)

	for {
		opt := utils.GetInputInt("\nTekan 1 untuk konfirmasi pembayaran\nTekan 2 untuk kembali ke menu utama\nSilakan pilih 1/2 : ")
		if opt == 1 {
			fmt.Println("Terima kasih sudah berbelanja")
			fmt.Println("[ !!! DIKEMBALIKAN KE MENU UTAMA !!! ]")
			module.AddTransactions()
			module.EmptyCartItem()
			time.Sleep(time.Second)
			return
		} else if opt == 2 {
			return
		} else {
			utils.InvalidInput()
		}
	}
}

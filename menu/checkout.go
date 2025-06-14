package menu

import (
	"fgo24-weekly-go/module"
	"fgo24-weekly-go/utils"
	"fmt"
	"time"
)

func CheckOut(cart *module.CartManager, transactions *module.TransactionManager) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Checkout :")

	items := cart.GetAll()
	if len(items) == 0 {
		fmt.Println("Belum ada item yang dimasukkan ke dalam keranjang!")
		fmt.Printf("\n[DIKEMBALIKAN KE MENU UTAMA]\n\n")
		time.Sleep(2 * time.Second)
		return
	}

	var totalPrice int = 0
	for i, item := range items {
		fmt.Printf("%d. %s - Rp.%d\n", i+1, item.Name, item.Price)
		totalPrice = totalPrice + item.Price
	}
	fmt.Printf("Total Pembayaran : Rp.%d\n", totalPrice)

	for {
		opt := utils.GetInputInt("\nTekan 1 untuk konfirmasi pembayaran\nTekan 2 untuk kembali ke menu utama\nSilakan pilih 1/2 : ")
		if opt == 1 {
			fmt.Println("\nTerima kasih sudah berbelanja")
			fmt.Println("[ !!! DIKEMBALIKAN KE MENU UTAMA !!! ]")
			transactions.Add(cart.GetAll())
			cart.EmptyCart()
			time.Sleep(time.Second)
			return
		} else if opt == 2 {
			return
		} else {
			utils.InvalidInput()
		}
	}
}

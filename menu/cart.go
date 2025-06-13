package menu

import (
	"fgo24-weekly-go/module"
	"fgo24-weekly-go/utils"
	"fmt"
	"time"
)

func ShoppingCart() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Keranjang :")
	if len(module.Cart) == 0 {
		fmt.Println("Belum ada item yang dimasukkan ke dalam keranjang!")
		fmt.Printf("\n[DIKEMBALIKAN KE MENU UTAMA]\n\n")
		time.Sleep(2 * time.Second)
		MenuCustomer()
	}
	var totalPrice int = 0
	for i, item := range module.Cart {
		fmt.Printf("%d. %s - Rp.%d\n", i+1, item.Name, item.Price)
		totalPrice = totalPrice + item.Price
	}
	fmt.Printf("Total Pembayaran : Rp.%d\n\n", totalPrice)
	for {
		if len(module.Cart) == 1 {
			fmt.Print("Pilih 1 untuk menghapus item dari Keranjang\n")
		} else {
			fmt.Printf("Pilih 1-%d untuk menghapus item dari Keranjang\n", len(module.Cart))
		}
		fmt.Printf("Pilih %d untuk Checkout\n", len(module.Cart)+1)
		fmt.Printf("Pilih %d untuk kembali ke menu utama\n", len(module.Cart)+2)
		opt := utils.GetInputInt(fmt.Sprintf("Silahkan pilih 1-%d : ", len(module.Cart)+2))
		if opt < 1 || opt > len(module.Cart)+2 {
			utils.InvalidInput()
		} else if opt == len(module.Cart)+1 {
			CheckOut()
		} else if opt == len(module.Cart)+2 {
			return
		} else {
			module.RemoveCartItem(opt)
			ShoppingCart()
		}
	}
}

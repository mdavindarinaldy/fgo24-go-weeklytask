package menu

import (
	"fgo24-weekly-go/module"
	"fgo24-weekly-go/utils"
	"fmt"
	"time"
)

func ShoppingCart(cart *module.CartManager) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Keranjang :")
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
	fmt.Printf("Total Pembayaran : Rp.%d\n\n", totalPrice)
	for {
		if len(items) == 1 {
			fmt.Print("Pilih 1 untuk menghapus item dari Keranjang\n")
		} else {
			fmt.Printf("Pilih 1-%d untuk menghapus item dari Keranjang\n", len(items))
		}
		fmt.Printf("Pilih %d untuk kembali ke menu utama\n", len(items)+1)
		opt := utils.GetInputInt(fmt.Sprintf("Silahkan pilih 1-%d : ", len(items)+1))
		if opt < 1 || opt > len(items)+1 {
			utils.InvalidInput()
		} else if opt == len(items)+1 {
			return
		} else {
			stringConfirm := fmt.Sprintf("\nApakah yakin ingin menghapus %s dari keranjang?\nPilih 1 untuk konfirmasi menghapus item, masukkan selain 1 untuk membatalkan penghapusan item: ", items[opt-1].Name)
			confirm := utils.GetInputInt(stringConfirm)
			if confirm == 1 {
				cart.Remove(opt)
				fmt.Println("\nItem berhasil dihapus!")
				fmt.Printf("\n[DIKEMBALIKAN KE MENU UTAMA]\n\n")
				time.Sleep(2 * time.Second)
				return
			} else {
				fmt.Println("\nPenghapusan item dibatalkan!")
				fmt.Printf("\n[DIKEMBALIKAN KE MENU UTAMA]\n\n")
				time.Sleep(2 * time.Second)
				return
			}
		}
	}
}

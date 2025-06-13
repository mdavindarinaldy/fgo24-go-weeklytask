package menu

import (
	"fgo24-weekly-go/module"
	"fgo24-weekly-go/utils"
	"fmt"
	"strings"
	"time"
)

func MenuAdmin() {
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("----------MENU ADMIN----------")
		fmt.Println("1. Tambah menu baru")
		fmt.Println("2. Hapus menu")
		fmt.Println("3. Rangkuman transaksi")
		fmt.Println("4. Logout")
		option := utils.GetInputInt("Silakan pilih menu [1-4] : ")
		if option == 1 {
			AddItemMenu()
		} else if option == 2 {
			fmt.Println("hapus menu")
		} else if option == 3 {
			fmt.Println("tampilkan transaksi")
		} else if option == 4 {
			return
		} else {
			utils.InvalidInput()
		}
	}
}

func AddItemMenu() {
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Menu Baru :")
		name := utils.GetInputString("Masukkan nama hidangan : ")
		price := utils.GetInputInt("Masukkan harga hidangan : ")
		origin := utils.GetInputString("Khas manakah hidangan tersebut (western/nusantara/japanese) : ")
		category := utils.GetInputString("Termasuk kategori apakah hidangan tersebut (food/drink/snack/dessert) : ")
		fmt.Println("\nPilih 1 untuk melanjutkan konfirmasi penambahan menu baru")
		fmt.Println("Pilih 2 untuk melanjutkan membatalkan penambahan menu baru dan kembali ke menu utama")
		opt := utils.GetInputInt("Pilih 1/2 : ")

		if opt == 2 {
			fmt.Println("[ !!! DIKEMBALIKAN KE MENU UTAMA !!! ]")
			return
		} else if opt == 1 {
			if module.CheckItem(name) {
				fmt.Printf("\n\nGagal menambahkan karena item dengan nama %s sudah ada di menu, silakan coba lagi!\n", name)
				time.Sleep(time.Second)
			} else if origin == "western" || origin == "nusantara" || origin == "japanese" || category == "food" || category == "drink" || category == "snack" || category == "dessert" || price > 0 {
				newItem := module.Item{
					Name:     name,
					Price:    price,
					Origin:   strings.ToLower(origin),
					Category: strings.ToLower(category),
				}
				module.AddItem(newItem)
				fmt.Printf("\n\n%s berhasil ditambahkan ke dalam menu\n", newItem.Name)
				time.Sleep(time.Second)
				return
			} else {
				utils.InvalidInput()
			}
		} else {
			utils.InvalidInput()
		}
	}
}

func RemoveItemMenu() {

}

func DisplayTransactions() {

}

func DisplayAllItems() {

}

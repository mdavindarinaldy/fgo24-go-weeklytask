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
		fmt.Println("4. Lihat semua item menu")
		fmt.Println("5. Logout")
		option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
		if option == 1 {
			AddItemMenu()
		} else if option == 2 {
			fmt.Println("hapus menu")
		} else if option == 3 {
			DisplayTransactions()
		} else if option == 4 {
			DisplayAllItems()
		} else if option == 5 {
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
	fmt.Print("\033[H\033[2J")
	if len(module.Transactions) == 0 {
		fmt.Println("Belum ada transaksi!")
		fmt.Printf("\n[DIKEMBALIKAN KE MENU UTAMA]\n\n")
		time.Sleep(2 * time.Second)
		MenuCustomer()
	}
	module.SortTransactions()
	fmt.Printf("Total Transaksi : %d\n\n", len(module.Transactions))
	fmt.Println("Rangkuman transaksi : ")
	for i, item := range module.SummaryTransactions {
		fmt.Printf("%d. %s: \n", i+1, item.Name)
		fmt.Printf("-- Dipesan pelanggan sebanyak %d kali\n", item.Order)
		fmt.Printf("-- Total pembelian %d kali\n", item.Price*item.Order)
	}
	opt := utils.GetInputInt("\nPilih 1 untuk kembali ke menu utama : ")
	if opt == 1 {
		return
	} else {
		utils.InvalidInput()
		DisplayTransactions()
	}
}

func DisplayAllItems() {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Total Item : %d\n\n", len(module.Items))
	fmt.Println("List semua item menu :")
	for i, item := range module.Items {
		fmt.Printf("%d. %s -- %s: Rp.%d\n", i+1, item.Name, item.Origin, item.Price)
	}
	opt := utils.GetInputInt("\nPilih 1 untuk kembali ke menu utama : ")
	if opt == 1 {
		return
	} else {
		utils.InvalidInput()
		DisplayAllItems()
	}
}

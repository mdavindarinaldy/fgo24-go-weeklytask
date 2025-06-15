package menu

import (
	"fgo24-weekly-go/module"
	"fgo24-weekly-go/utils"
	"fmt"
	"math"
	"strings"
	"time"
)

func MenuAdmin(menu *module.Menu, transactions *module.TransactionManager) {
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
			AddItemMenu(menu)
		} else if option == 2 {
			RemoveItemMenu(menu)
		} else if option == 3 {
			DisplayTransactions(transactions)
		} else if option == 4 {
			DisplayAllItems(menu, 1)
		} else if option == 5 {
			return
		} else {
			utils.InvalidInput()
		}
	}
}

func AddItemMenu(menu *module.Menu) {
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
			if menu.Check(name) {
				fmt.Printf("\n\nGagal menambahkan karena item dengan nama %s sudah ada di menu, silakan coba lagi!\n", name)
				time.Sleep(time.Second)
			} else if origin == "western" || origin == "nusantara" || origin == "japanese" || category == "food" || category == "drink" || category == "snack" || category == "dessert" || price > 0 {
				newItem := module.Item{
					Name:     name,
					Price:    price,
					Origin:   strings.ToLower(origin),
					Category: strings.ToLower(category),
				}
				menu.Add(newItem)
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

func RemoveItemMenu(menu *module.Menu) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Hapus Item :")
	items := menu.GetAll()
	for i, item := range items {
		fmt.Printf("%d. %s - %s - Rp.%d\n", i+1, item.Name, item.Category, item.Price)
	}
	for {
		if len(items) == 1 {
			fmt.Print("Item menu tersisa satu dan tidak boleh dihapus agar menu tidak kosong! Silahkan tambahkan item terlebih dahulu untuk menghapus\n")
			fmt.Printf("\n[DIKEMBALIKAN KE MENU UTAMA]\n\n")
			time.Sleep(2 * time.Second)
			return
		} else {
			fmt.Printf("\nPilih 1-%d untuk menghapus item dari Keranjang\n", len(items))
		}
		fmt.Printf("Pilih %d untuk kembali ke menu utama\n", len(items)+1)
		opt := utils.GetInputInt(fmt.Sprintf("Silahkan pilih 1-%d : ", len(items)+1))
		if opt < 1 || opt > len(items)+1 {
			utils.InvalidInput()
		} else if opt == len(items)+1 {
			return
		} else {
			stringConfirm := fmt.Sprintf("\nApakah yakin ingin menghapus %s dari menu?\nPilih 1 untuk konfirmasi menghapus item, masukkan selain 1 untuk membatalkan penghapusan item: ", items[opt-1].Name)
			confirm := utils.GetInputInt(stringConfirm)
			if confirm == 1 {
				menu.Remove(opt)
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

func DisplayTransactions(transactions *module.TransactionManager) {
	fmt.Print("\033[H\033[2J")
	items := transactions.GetAll()
	if len(items) == 0 {
		fmt.Println("Belum ada transaksi!")
		fmt.Printf("\n[DIKEMBALIKAN KE MENU UTAMA]\n\n")
		time.Sleep(2 * time.Second)
		return
	}
	fmt.Printf("Total Transaksi : %d\n\n", len(items))
	fmt.Println("Rangkuman transaksi : ")
	transactions.SortTransaction()
	for i, item := range transactions.GetSummary() {
		fmt.Printf("%d. %s: \n", i+1, item.Name)
		fmt.Printf("-- Dipesan pelanggan sebanyak %d kali\n", item.Order)
		fmt.Printf("-- Total pembelian %d kali\n", item.Price*item.Order)
	}
	opt := utils.GetInputInt("\nPilih 1 untuk kembali ke menu utama : ")
	if opt == 1 {
		return
	} else {
		utils.InvalidInput()
		DisplayTransactions(transactions)
	}
}

func DisplayAllItems(menu *module.Menu, page int) {
	items := menu.GetAll()
	itemsPerPage := 5
	startIndex := (page - 1) * itemsPerPage
	endIndex := startIndex + itemsPerPage
	totalPages := int(math.Ceil(float64(len(items)) / float64(itemsPerPage)))

	if endIndex > len(items) {
		endIndex = len(items)
	}

	if startIndex >= 0 && startIndex < len(items) {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Total Page: %d --- Total Item: %d --- Page Saat Ini: %d\n\n", totalPages, len(items), page)
		fmt.Println("LIST SEMUA ITEM")
		for i, item := range items[startIndex:endIndex] {
			fmt.Printf("%d. %s -- %s: Rp.%d\n", startIndex+i+1, item.Name, item.Origin, item.Price)
		}
		if startIndex == 0 {
			fmt.Println("\nPilih 1 untuk ke halaman berikutnya")
			fmt.Println("Pilih 2 untuk kembali ke menu utama")
			opt := utils.GetInputInt("\nSilakan pilih menu dengan memilih angka : ")
			if opt == 2 {
				return
			} else if opt == 1 {
				DisplayAllItems(menu, page+1)
			} else {
				utils.InvalidInput()
				DisplayAllItems(menu, page)
			}
		} else if endIndex == len(items) {
			fmt.Println("\nPilih 1 untuk ke halaman sebelumnya")
			fmt.Println("Pilih 2 untuk kembali ke menu utama")
			opt := utils.GetInputInt("\nSilakan pilih menu dengan memilih angka : ")
			if opt == 2 {
				return
			} else if opt == 1 {
				DisplayAllItems(menu, page-1)
			} else {
				utils.InvalidInput()
				DisplayAllItems(menu, page)
			}
		} else {
			fmt.Println("\nPilih 1 untuk ke halaman sebelumnya")
			fmt.Println("Pilih 2 untuk ke halaman berikutnya")
			fmt.Println("Pilih 3 untuk kembali ke menu utama")
			opt := utils.GetInputInt("\nSilakan pilih menu dengan memilih angka : ")
			if opt == 3 {
				return
			} else if opt == 1 {
				DisplayAllItems(menu, page-1)
			} else if opt == 2 {
				DisplayAllItems(menu, page+1)
			} else {
				utils.InvalidInput()
				DisplayAllItems(menu, page)
			}
		}
	}
}

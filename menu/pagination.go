package menu

import (
	"fgo24-weekly-go/module"
	"fgo24-weekly-go/utils"
	"fmt"
	"math"
)

func Pagination(items []module.Item, page int, cart *module.CartManager) {
	itemsPerPage := 5
	startIndex := (page - 1) * itemsPerPage
	endIndex := startIndex + itemsPerPage
	totalPages := int(math.Ceil(float64(len(items)) / float64(itemsPerPage)))

	if endIndex > len(items) {
		endIndex = len(items)
	}

	if startIndex >= 0 && startIndex < len(items) {
		fmt.Print("\033[H\033[2J")
		fmt.Println("LIST MENU ITEM")
		fmt.Printf("Total Page: %d --- Page Saat Ini: %d\n\n", totalPages, page)
		for i, item := range items[startIndex:endIndex] {
			fmt.Printf("%d. %s : Rp.%d\n", startIndex+i+1, item.Name, item.Price)
		}
		fmt.Printf("\nPilih %d-%d untuk menambahkan item ke keranjang\n", startIndex+1, endIndex)
		if startIndex == 0 {
			fmt.Printf("Pilih %d untuk ke halaman berikutnya\n", endIndex+1)
			fmt.Printf("Pilih %d untuk kembali ke menu utama\n", endIndex+2)
			option := utils.GetInputInt("\nSilakan pilih menu dengan memilih angka : ")
			if option < 1 || option > endIndex+2 {
				utils.InvalidInput()
				Pagination(items, page, cart)
			} else if option == endIndex+1 {
				Pagination(items, page+1, cart)
			} else if option == endIndex+2 {
				return
			} else {
				cart.Add(items[option-1])
			}
		} else if endIndex == len(items) {
			fmt.Printf("Pilih %d untuk ke halaman sebelumnya\n", endIndex+1)
			fmt.Printf("Pilih %d untuk kembali ke menu utama\n", endIndex+2)
			option := utils.GetInputInt("\nSilakan pilih menu dengan memilih angka : ")
			if option < 1 || option > endIndex+2 {
				utils.InvalidInput()
				Pagination(items, page, cart)
			} else if option == endIndex+1 {
				Pagination(items, page-1, cart)
			} else if option == endIndex+2 {
				return
			} else {
				cart.Add(items[option-1])
			}
		} else {
			fmt.Printf("Pilih %d untuk ke halaman sebelumnya\n", endIndex+1)
			fmt.Printf("Pilih %d untuk ke halaman berikutnya\n", endIndex+2)
			fmt.Printf("Pilih %d untuk kembali ke menu utama\n", endIndex+3)
			option := utils.GetInputInt("\nSilakan pilih menu dengan memilih angka : ")
			if option < 1 || option > endIndex+3 {
				utils.InvalidInput()
				Pagination(items, page, cart)
			} else if option == endIndex+1 {
				Pagination(items, page-1, cart)
			} else if option == endIndex+2 {
				Pagination(items, page+1, cart)
			} else if option == endIndex+3 {
				return
			} else {
				cart.Add(items[option-1])
			}
		}
	}
}

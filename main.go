package main

import (
	"fgo24-weekly-go/menu"
	"fgo24-weekly-go/utils"
	"fmt"
	"os"
)

func main() {
	for {
		menu.MainMenu()
		option := utils.GetInputInt("Pilih Menu 1-4: ")
		if option == 1 {
			menu.Menu()
		} else if option == 2 {
			menu.ShoppingCart()
		} else if option == 3 {
			menu.CheckOut()
		} else if option == 4 {
			fmt.Println("Terima kasih sudah berkunjung!")
			os.Exit(0)
		}
	}
}

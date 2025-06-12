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
		option := utils.GetInputInt("Pilih Menu 1-3: ")
		if option == 1 {
			menu.MenuCustomer()
		} else if option == 2 {
			fmt.Println("Menu Admin")
		} else if option == 3 {
			fmt.Println("Program Diakhiri!")
			os.Exit(0)
		} else {
			utils.InvalidInput()
		}
	}
}

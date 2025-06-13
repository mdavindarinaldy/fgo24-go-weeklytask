package main

import (
	"fgo24-weekly-go/menu"
	"fgo24-weekly-go/utils"
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Print("\033[H\033[2J") // clear screen
		menu.MainMenu()
		option := utils.GetInputInt("Pilih Menu 1-3: ")
		if option == 1 {
			menu.MenuCustomer()
		} else if option == 2 {
			menu.MenuAdmin()
		} else if option == 3 {
			fmt.Print("\n[ !!! PROGRAM DIAKHIRI !!! ]\n\n")
			time.Sleep(2 * time.Second)
			os.Exit(0)
		} else {
			utils.InvalidInput()
		}
	}
}

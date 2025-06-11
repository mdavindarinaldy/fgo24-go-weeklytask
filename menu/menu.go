package menu

import (
	"fgo24-weekly-go/utils"
	"fmt"
)

func MainMenu() {
	fmt.Println("Selamat datang di Warung!")
	fmt.Println("1. Lihat menu")
	fmt.Println("2. Lihat keranjang")
	fmt.Println("3. Checkout")
	fmt.Println("4. Exit Program")
}

func Menu() {
	fmt.Println("Menu :")
	fmt.Println("1. Menu Makanan")
	fmt.Println("2. Menu Minuman")
	fmt.Println("3. Menu Makanan Ringan")
	fmt.Println("4. Menu Desserts")
	fmt.Println("5. Kembali ke Menu Utama")
	option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
	if option == 1 {
		FoodsMenu()
	}
	if option == 2 {
		DrinksMenu()
	}
	if option == 3 {
		SnacksMenu()
	}
	if option == 4 {
		DessertsMenu()
	}
	if option == 5 {
		return
	}
}

func FoodsMenu() {
	fmt.Println("Menu Makanan :")
	for i, item := range FoodItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.name, item.price)
	}
	fmt.Println("5. Kembali ke Menu Utama")
	option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
	if option < 1 || option > 5 {
		fmt.Println("Input tidak valid! Silakan coba lagi")
		FoodsMenu()
	} else if option == 5 {
		return
	} else {
		AddItem(FoodItems[option-1])
	}
}

func DrinksMenu() {
	fmt.Println("Menu Minuman :")
	for i, item := range DrinkItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.name, item.price)
	}
	fmt.Println("5. Kembali ke Menu Utama")
	option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
	if option < 1 || option > 5 {
		fmt.Println("Input tidak valid! Silakan coba lagi")
		FoodsMenu()
	} else if option == 5 {
		return
	} else {
		AddItem(DrinkItems[option-1])
	}
}

func SnacksMenu() {
	fmt.Println("Menu Snacks :")
	for i, item := range SnackItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.name, item.price)
	}
	fmt.Println("5. Kembali ke Menu Utama")
	option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
	if option < 1 || option > 5 {
		fmt.Println("Input tidak valid! Silakan coba lagi")
		FoodsMenu()
	} else if option == 5 {
		return
	} else {
		AddItem(SnackItems[option-1])
	}
}

func DessertsMenu() {
	fmt.Println("Menu Desserts :")
	for i, item := range DessertItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.name, item.price)
	}
	fmt.Println("5. Kembali ke Menu Utama")
	option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
	if option < 1 || option > 5 {
		fmt.Println("Input tidak valid! Silakan coba lagi")
		FoodsMenu()
	} else if option == 5 {
		return
	} else {
		AddItem(DessertItems[option-1])
	}
}

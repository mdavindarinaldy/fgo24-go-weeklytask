package menu

import (
	"fgo24-weekly-go/module"
	"fgo24-weekly-go/utils"
	"fmt"
	"sync"
)

func MenuCustomer() {
	for {
		wait := sync.WaitGroup{}
		wait.Add(1)
		go module.EmptyItem(&wait)
		wait.Wait()
		fmt.Print("\n----------MENU UTAMA----------\n")
		fmt.Println("1. Lihat menu berdasarkan kategori")
		fmt.Println("2. Lihat menu berdasarkan tipe hidangan")
		fmt.Println("3. Lihat menu berdasarkan harga terrendah")
		fmt.Println("4. Lihat menu berdasarkan harga tertinggi")
		fmt.Println("5. Lihat keranjang")
		fmt.Println("6. Checkout")
		fmt.Println("7. Logout")
		option := utils.GetInputInt("Silakan pilih menu [1-7] : ")
		if option == 1 {
			MenuByCategory()
		} else if option == 2 {
			fmt.Println("Menu Tipe")
		} else if option == 3 {
			fmt.Println("Menu Harga Terrendah")
		} else if option == 4 {
			fmt.Println("Menu Harga Tertinggi")
		} else if option == 5 {
			ShoppingCart()
		} else if option == 6 {
			CheckOut()
		} else if option == 7 {
			return
		} else {
			utils.InvalidInput()
		}
	}
}

func MenuByCategory() {
	wait := sync.WaitGroup{}
	wait.Add(4)
	go module.FoodItem(&wait)
	go module.DrinkItem(&wait)
	go module.SnackItem(&wait)
	go module.DessertItem(&wait)
	wait.Wait()

	for {
		fmt.Println("Menu Berdasarkan Kategori :")
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
}

func FoodsMenu() {
	fmt.Println("Menu Makanan :")
	for i, item := range module.FoodItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Println("5. Kembali ke Menu Sebelumnya")
	option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
	if option < 1 || option > 5 {
		fmt.Println("Input tidak valid! Silakan coba lagi")
		FoodsMenu()
	} else if option == 5 {
		return
	} else {
		module.AddCartItem(module.FoodItems[option-1])
	}
}

func DrinksMenu() {
	fmt.Println("Menu Minuman :")
	for i, item := range module.DrinkItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Println("5. Kembali ke Menu Utama")
	option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
	if option < 1 || option > 5 {
		fmt.Println("Input tidak valid! Silakan coba lagi")
		FoodsMenu()
	} else if option == 5 {
		return
	} else {
		module.AddCartItem(module.DrinkItems[option-1])
	}
}

func SnacksMenu() {
	fmt.Println("Menu Snacks :")
	for i, item := range module.SnackItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Println("5. Kembali ke Menu Utama")
	option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
	if option < 1 || option > 5 {
		fmt.Println("Input tidak valid! Silakan coba lagi")
		FoodsMenu()
	} else if option == 5 {
		return
	} else {
		module.AddCartItem(module.SnackItems[option-1])
	}
}

func DessertsMenu() {
	fmt.Println("Menu Desserts :")
	for i, item := range module.DessertItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Println("5. Kembali ke Menu Utama")
	option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
	if option < 1 || option > 5 {
		fmt.Println("Input tidak valid! Silakan coba lagi")
		FoodsMenu()
	} else if option == 5 {
		return
	} else {
		module.AddCartItem(module.DessertItems[option-1])
	}
}

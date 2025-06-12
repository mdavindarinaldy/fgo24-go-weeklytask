package menu

import (
	"fgo24-weekly-go/module"
	"fgo24-weekly-go/utils"
	"fmt"
	"sync"
)

func MenuCustomer() {
	for {
		fmt.Print("\033[H\033[2J")
		wait := sync.WaitGroup{}
		wait.Add(2)
		go module.EmptyItemByCategory(&wait)
		go module.EmptyItemByOrigin(&wait)
		wait.Wait()
		fmt.Println("----------MENU UTAMA----------")
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
			MenuByOrigin()
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
		fmt.Print("\033[H\033[2J")
		fmt.Println("Menu Berdasarkan Kategori :")
		fmt.Println("1. Menu Makanan")
		fmt.Println("2. Menu Minuman")
		fmt.Println("3. Menu Makanan Ringan")
		fmt.Println("4. Menu Desserts")
		fmt.Println("5. Kembali ke Menu Utama")
		option := utils.GetInputInt("Silakan pilih menu [1-5] : ")
		if option == 1 {
			FoodsMenu()
		} else if option == 2 {
			DrinksMenu()
		} else if option == 3 {
			SnacksMenu()
		} else if option == 4 {
			DessertsMenu()
		} else if option == 5 {
			return
		} else {
			utils.InvalidInput()
		}
	}
}

func MenuByOrigin() {
	wait := sync.WaitGroup{}
	wait.Add(3)
	go module.NusantaraItem(&wait)
	go module.WesternItem(&wait)
	go module.JapaneseItem(&wait)
	wait.Wait()

	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Menu Berdasarkan Tipe Hidangan :")
		fmt.Println("1. Hidangan Nusantara")
		fmt.Println("2. Hidangan Western")
		fmt.Println("3. Hidangan Japanese")
		fmt.Println("4. Kembali ke Menu Utama")
		option := utils.GetInputInt("Silakan pilih menu [1-4] : ")
		if option == 1 {
			NusantaraMenu()
		} else if option == 2 {
			WesternMenu()
		} else if option == 3 {
			JapaneseMenu()
		} else if option == 4 {
			return
		} else {
			utils.InvalidInput()
		}
	}
}

func FoodsMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Menu Makanan :")
	len := len(module.FoodItems)
	for i, item := range module.FoodItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		FoodsMenu()
	} else if option == len+1 {
		return
	} else {
		module.AddCartItem(module.FoodItems[option-1])
	}
}

func DrinksMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Menu Minuman :")
	len := len(module.DrinkItems)
	for i, item := range module.DrinkItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		fmt.Println("Input tidak valid! Silakan coba lagi")
		DrinksMenu()
	} else if option == len+1 {
		return
	} else {
		module.AddCartItem(module.DrinkItems[option-1])
	}
}

func SnacksMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Menu Snacks :")
	len := len(module.SnackItems)
	for i, item := range module.SnackItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		SnacksMenu()
	} else if option == len+1 {
		return
	} else {
		module.AddCartItem(module.SnackItems[option-1])
	}
}

func DessertsMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Menu Desserts :")
	len := len(module.DessertItems)
	for i, item := range module.DessertItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		DessertsMenu()
	} else if option == len+1 {
		return
	} else {
		module.AddCartItem(module.DessertItems[option-1])
	}
}

func NusantaraMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Hidangan Nusantara :")
	len := len(module.NusantaraItems)
	for i, item := range module.NusantaraItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		NusantaraMenu()
	} else if option == len+1 {
		return
	} else {
		module.AddCartItem(module.NusantaraItems[option-1])
	}
}

func WesternMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Hidangan Nusantara :")
	len := len(module.WesternItems)
	for i, item := range module.WesternItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		WesternMenu()
	} else if option == len+1 {
		return
	} else {
		module.AddCartItem(module.WesternItems[option-1])
	}
}

func JapaneseMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Hidangan Nusantara :")
	len := len(module.JapaneseItems)
	for i, item := range module.JapaneseItems {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		JapaneseMenu()
	} else if option == len+1 {
		return
	} else {
		module.AddCartItem(module.JapaneseItems[option-1])
	}
}

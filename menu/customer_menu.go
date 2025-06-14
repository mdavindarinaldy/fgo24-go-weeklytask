package menu

import (
	"fgo24-weekly-go/module"
	"fgo24-weekly-go/utils"
	"fmt"
	"sync"
)

func MenuCustomer(menu *module.Menu, transactions *module.TransactionManager) {
	cart := module.NewCartManager()
	sortedMenu := module.NewSortedMenu(menu)
	menuByOrigin := module.NewMenuByOrigin()
	menuByCategory := module.NewMenuByCategory()

	for {
		fmt.Print("\033[H\033[2J")
		wait := sync.WaitGroup{}
		wait.Add(3)
		go menuByCategory.EmptyItems(&wait)
		go menuByOrigin.EmptyItems(&wait)
		go sortedMenu.ResetSortedItem(&wait, menu)
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
			MenuByCategory(menu, cart, menuByCategory)
		} else if option == 2 {
			MenuByOrigin(menu, cart, menuByOrigin)
		} else if option == 3 {
			SortAscendMenu(cart, sortedMenu)
		} else if option == 4 {
			SortDescendMenu(cart, sortedMenu)
		} else if option == 5 {
			ShoppingCart(cart)
		} else if option == 6 {
			CheckOut(cart, transactions)
		} else if option == 7 {
			return
		} else {
			utils.InvalidInput()
		}
	}
}

func MenuByCategory(menu *module.Menu, cart *module.CartManager, menuByCategory *module.MenuByCategory) {
	wait := sync.WaitGroup{}
	wait.Add(4)
	go menuByCategory.FilterFood(&wait, menu)
	go menuByCategory.FilterDrink(&wait, menu)
	go menuByCategory.FilterSnack(&wait, menu)
	go menuByCategory.FilterDessert(&wait, menu)
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
			FoodsMenu(cart, menuByCategory)
		} else if option == 2 {
			DrinksMenu(cart, menuByCategory)
		} else if option == 3 {
			SnacksMenu(cart, menuByCategory)
		} else if option == 4 {
			DessertsMenu(cart, menuByCategory)
		} else if option == 5 {
			return
		} else {
			utils.InvalidInput()
		}
	}
}

func MenuByOrigin(menu *module.Menu, cart *module.CartManager, menuByOrigin *module.MenuByOrigin) {
	wait := sync.WaitGroup{}
	wait.Add(3)
	go menuByOrigin.FilterNusantara(&wait, menu)
	go menuByOrigin.FilterWestern(&wait, menu)
	go menuByOrigin.FilterJapanese(&wait, menu)
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
			NusantaraMenu(cart, menuByOrigin)
		} else if option == 2 {
			WesternMenu(cart, menuByOrigin)
		} else if option == 3 {
			JapaneseMenu(cart, menuByOrigin)
		} else if option == 4 {
			return
		} else {
			utils.InvalidInput()
		}
	}
}

func FoodsMenu(cart *module.CartManager, menuByCategory *module.MenuByCategory) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Menu Makanan :")
	items := menuByCategory.GetFoodItems()
	len := len(items)
	for i, item := range items {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		FoodsMenu(cart, menuByCategory)
	} else if option == len+1 {
		return
	} else {
		cart.Add(items[option-1])
	}
}

func DrinksMenu(cart *module.CartManager, menuByCategory *module.MenuByCategory) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Menu Minuman :")
	items := menuByCategory.GetDrinkItems()
	len := len(items)
	for i, item := range items {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		fmt.Println("Input tidak valid! Silakan coba lagi")
		DrinksMenu(cart, menuByCategory)
	} else if option == len+1 {
		return
	} else {
		cart.Add(items[option-1])
	}
}

func SnacksMenu(cart *module.CartManager, menuByCategory *module.MenuByCategory) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Menu Snacks :")
	items := menuByCategory.GetSnackItems()
	len := len(items)
	for i, item := range items {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		SnacksMenu(cart, menuByCategory)
	} else if option == len+1 {
		return
	} else {
		cart.Add(items[option-1])
	}
}

func DessertsMenu(cart *module.CartManager, menuByCategory *module.MenuByCategory) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Menu Desserts :")
	items := menuByCategory.GetSnackItems()
	len := len(items)
	for i, item := range items {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		DessertsMenu(cart, menuByCategory)
	} else if option == len+1 {
		return
	} else {
		cart.Add(items[option-1])
	}
}

func NusantaraMenu(cart *module.CartManager, menuByOrigin *module.MenuByOrigin) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Hidangan Nusantara :")
	items := menuByOrigin.GetNusantaraItems()
	len := len(items)
	for i, item := range items {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		NusantaraMenu(cart, menuByOrigin)
	} else if option == len+1 {
		return
	} else {
		cart.Add(items[option-1])
	}
}

func WesternMenu(cart *module.CartManager, menuByOrigin *module.MenuByOrigin) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Hidangan Nusantara :")
	items := menuByOrigin.GetWesternItems()
	len := len(items)
	for i, item := range items {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		WesternMenu(cart, menuByOrigin)
	} else if option == len+1 {
		return
	} else {
		cart.Add(items[option-1])
	}
}

func JapaneseMenu(cart *module.CartManager, menuByOrigin *module.MenuByOrigin) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Hidangan Nusantara :")
	items := menuByOrigin.GetJapaneseItems()
	len := len(items)
	for i, item := range items {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		JapaneseMenu(cart, menuByOrigin)
	} else if option == len+1 {
		return
	} else {
		cart.Add(items[option-1])
	}
}

func SortAscendMenu(cart *module.CartManager, sortedMenu *module.SortedMenu) {
	sortedMenu.SortAscend()
	fmt.Print("\033[H\033[2J")
	fmt.Println("List menu dari harga terrendah ke tertinggi : ")
	items := sortedMenu.GetSortedItems()
	len := len(items)
	for i, item := range items {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		SortAscendMenu(cart, sortedMenu)
	} else if option == len+1 {
		return
	} else {
		cart.Add(items[option-1])
	}
}

func SortDescendMenu(cart *module.CartManager, sortedMenu *module.SortedMenu) {
	sortedMenu.SortDescend()
	fmt.Print("\033[H\033[2J")
	fmt.Println("List menu dari harga tertinggi ke terrendah : ")
	items := sortedMenu.GetSortedItems()
	len := len(items)
	for i, item := range items {
		fmt.Printf("%d. %s : Rp.%d\n", i+1, item.Name, item.Price)
	}
	fmt.Printf("%d. Kembali ke Menu Sebelumnya\n", len+1)
	option := utils.GetInputInt("Silakan pilih menu dengan memilih angka : ")
	if option < 1 || option > len+1 {
		utils.InvalidInput()
		SortDescendMenu(cart, sortedMenu)
	} else if option == len+1 {
		return
	} else {
		cart.Add(items[option-1])
	}
}

package module

import (
	"sort"
	"sync"
)

type SortedMenu struct {
	items []Item
}

func NewSortedMenu(menu *Menu) *SortedMenu {
	return &SortedMenu{
		items: menu.GetAll(),
	}
}

func (sm *SortedMenu) SortAscend() {
	sort.SliceStable(sm.items, func(i, j int) bool {
		return sm.items[i].Price < sm.items[j].Price
	})
}

func (sm *SortedMenu) SortDescend() {
	sort.SliceStable(sm.items, func(i, j int) bool {
		return sm.items[i].Price > sm.items[j].Price
	})
}

func (sm *SortedMenu) ResetSortedItem(wg *sync.WaitGroup, menu *Menu) {
	defer wg.Done()
	sm.items = menu.GetAll()
}

func (sm *SortedMenu) GetSortedItems() []Item {
	return sm.items
}

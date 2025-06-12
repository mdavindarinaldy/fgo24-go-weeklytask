package module

import (
	"sort"
	"sync"
)

var SortedItem []Item = Items

func SortAscend() {
	sort.SliceStable(SortedItem, func(i, j int) bool {
		return SortedItem[i].Price < SortedItem[j].Price
	})
}

func SortDescend() {
	sort.SliceStable(SortedItem, func(i, j int) bool {
		return SortedItem[i].Price > SortedItem[j].Price
	})
}

func ResetSortedItem(wg *sync.WaitGroup) {
	defer wg.Done()
	SortedItem = Items
}

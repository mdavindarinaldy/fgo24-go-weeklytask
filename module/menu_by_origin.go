package module

import "sync"

type MenuByOrigin struct {
	nusantaraItems []Item
	westernItems   []Item
	japaneseItems  []Item
}

func NewMenuByOrigin() *MenuByOrigin {
	return &MenuByOrigin{
		nusantaraItems: []Item{},
		westernItems:   []Item{},
		japaneseItems:  []Item{},
	}
}

func (mo *MenuByOrigin) FilterNusantara(wg *sync.WaitGroup, menu *Menu) {
	defer wg.Done()
	mo.nusantaraItems = []Item{}
	for _, item := range menu.GetAll() {
		if item.Origin == "nusantara" {
			mo.nusantaraItems = append(mo.nusantaraItems, item)
		}
	}
}

func (mo *MenuByOrigin) FilterWestern(wg *sync.WaitGroup, menu *Menu) {
	defer wg.Done()
	mo.westernItems = []Item{}
	for _, item := range menu.GetAll() {
		if item.Origin == "western" {
			mo.westernItems = append(mo.westernItems, item)
		}
	}
}

func (mo *MenuByOrigin) FilterJapanese(wg *sync.WaitGroup, menu *Menu) {
	defer wg.Done()
	mo.japaneseItems = []Item{}
	for _, item := range menu.GetAll() {
		if item.Origin == "japanese" {
			mo.japaneseItems = append(mo.japaneseItems, item)
		}
	}
}

// CHECK INI JUGAAA
func (mo *MenuByOrigin) EmptyItems(wg *sync.WaitGroup) {
	defer wg.Done()
	mo.nusantaraItems = []Item{}
	mo.westernItems = []Item{}
	mo.japaneseItems = []Item{}
}

func (mo *MenuByOrigin) GetNusantaraItems() []Item {
	return mo.nusantaraItems
}

func (mo *MenuByOrigin) GetWesternItems() []Item {
	return mo.westernItems
}

func (mo *MenuByOrigin) GetJapaneseItems() []Item {
	return mo.japaneseItems
}

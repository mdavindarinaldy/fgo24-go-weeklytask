package module

import "strings"

type ItemManager interface {
	Add(Item)
	Remove(int)
	Check(string) bool
	GetAll() []Item
}

type Item struct {
	Name     string
	Price    int
	Origin   string
	Category string
}

type Menu struct {
	items []Item
}

func NewMenu() *Menu {
	return &Menu{
		items: []Item{
			{
				Name:     "Nasi Goreng",
				Price:    15000,
				Origin:   "nusantara",
				Category: "food",
			}, {
				Name:     "Nasi Uduk",
				Price:    10000,
				Origin:   "nusantara",
				Category: "food",
			}, {
				Name:     "Fish and Chips",
				Price:    18000,
				Origin:   "western",
				Category: "food",
			}, {
				Name:     "Sushi",
				Price:    20000,
				Origin:   "japanese",
				Category: "food",
			}, {
				Name:     "Es Dawet",
				Price:    9000,
				Origin:   "nusantara",
				Category: "drink",
			}, {
				Name:     "Cherry Blossom Tea",
				Price:    5000,
				Origin:   "japanese",
				Category: "drink",
			}, {
				Name:     "Kombucha",
				Price:    8000,
				Origin:   "japanese",
				Category: "drink",
			}, {
				Name:     "Jasmine Tea in Affogato Latte",
				Price:    13000,
				Origin:   "western",
				Category: "drink",
			}, {
				Name:     "Mochi",
				Price:    9000,
				Origin:   "japanese",
				Category: "snack",
			}, {
				Name:     "Dango",
				Price:    10000,
				Origin:   "japanese",
				Category: "snack",
			}, {
				Name:     "Pisang Goreng",
				Price:    5000,
				Origin:   "nusantara",
				Category: "snack",
			}, {
				Name:     "Serabi",
				Price:    7500,
				Origin:   "nusantara",
				Category: "snack",
			}, {
				Name:     "Cheesecake",
				Price:    20000,
				Origin:   "western",
				Category: "dessert",
			}, {
				Name:     "Red Velvet Cake",
				Price:    22000,
				Origin:   "western",
				Category: "dessert",
			}, {
				Name:     "Creme Brulee",
				Price:    18000,
				Origin:   "western",
				Category: "dessert",
			}, {
				Name:     "Bubur Kacang Ijo",
				Price:    10000,
				Origin:   "nusantara",
				Category: "dessert",
			},
		},
	}
}

func (m *Menu) Add(item Item) {
	m.items = append(m.items, item)
}

func (m *Menu) Remove(index int) {
	if index == 1 {
		m.items = m.items[1:]
	} else if index == len(m.items) {
		m.items = m.items[:index-1]
	} else {
		m.items = append(m.items[:index-1], m.items[index:]...)
	}
}

func (m *Menu) Check(name string) bool {
	for _, item := range m.items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(name)) {
			return true
		}
	}
	return false
}

func (m *Menu) GetAll() []Item {
	return m.items
}

func (m *Menu) Search(name string) []Item {
	found := []Item{}
	for _, item := range m.items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(name)) {
			found = append(found, item)
		}
	}
	return found
}

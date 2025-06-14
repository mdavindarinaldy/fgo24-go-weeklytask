package module

import (
	"fmt"
	"strings"
	"time"
)

type CartManager struct {
	items []Item
}

func NewCartManager() *CartManager {
	return &CartManager{
		items: []Item{},
	}
}

func (c *CartManager) Add(item Item) {
	c.items = append(c.items, item)
	fmt.Printf("\n%s berhasil ditambahkan ke dalam keranjang!\n\n", item.Name)
	time.Sleep(time.Second)
}

func (c *CartManager) Remove(index int) {
	if index == 1 {
		c.items = c.items[1:]
	} else if index == len(c.items) {
		c.items = c.items[:index-1]
	} else {
		c.items = append(c.items[:index-1], c.items[index:]...)
	}
}

func (c *CartManager) Check(name string) bool {
	for _, item := range c.items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(name)) {
			return true
		}
	}
	return false
}

func (c *CartManager) GetAll() []Item {
	return c.items
}

func (c *CartManager) EmptyCart() {
	c.items = []Item{}
}

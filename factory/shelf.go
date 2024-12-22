package factory

import (
	"sync"
)

type Shelf struct {
	Name        string
	lock        *sync.Mutex
	Temperature string
	Capacity    int
	FoodOnShelf int
	orders      map[string]bool
}

func NewShelf(name string, temperature string, capacity int) *Shelf {
	orders := make(map[string]bool)

	shelf := &Shelf{
		Name:        name,
		Temperature: temperature,
		Capacity:    capacity,
		FoodOnShelf: 0,
		orders:      orders,
	}

	return shelf
}

func (s *Shelf) Register(order Order) {
	s.orders[order.Item.Id] = true
	s.FoodOnShelf += 1
}

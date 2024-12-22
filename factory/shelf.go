package factory

import (
	"sync"
)

type Shelf struct {
	Name        string
	lock        sync.Mutex
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
	s.lock.Lock()
	s.orders[order.Item.Id] = true
	s.FoodOnShelf += 1
	s.lock.Unlock()
}

func (s *Shelf) HasCapacity() bool {
	return (s.Capacity > s.FoodOnShelf)
}

func (s *Shelf) Contains(order Order) bool {
	orders := s.orders
	id := order.Item.Id

	_, exists := orders[id]

	return exists
}

func (s *Shelf) Remove(order Order) {
	s.lock.Lock()
	delete(s.orders, order.Item.Id)
	s.FoodOnShelf -= 1
	s.lock.Unlock()
}

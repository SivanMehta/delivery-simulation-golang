package factory

import (
	"sync"
)

type Shelf struct {
	Name        string
	Lock        sync.Mutex
	Temperature string
	Capacity    int
	FoodOnShelf int
	Orders      map[string]Item
}

func NewShelf(name string, temperature string, capacity int) *Shelf {
	orders := make(map[string]Item)

	shelf := &Shelf{
		Name:        name,
		Temperature: temperature,
		Capacity:    capacity,
		FoodOnShelf: 0,
		Orders:      orders,
	}

	return shelf
}

func (s *Shelf) Register(order Order) {
	s.Lock.Lock()
	s.Orders[order.Item.Id] = order.Item
	s.FoodOnShelf += 1
	s.Lock.Unlock()
}

func (s *Shelf) HasCapacity() bool {
	return (s.Capacity > s.FoodOnShelf)
}

func (s *Shelf) Contains(order Order) bool {
	orders := s.Orders
	id := order.Item.Id

	_, exists := orders[id]

	return exists
}

func (s *Shelf) Remove(order Order) {
	s.Lock.Lock()
	delete(s.Orders, order.Item.Id)
	s.FoodOnShelf -= 1
	s.Lock.Unlock()
}

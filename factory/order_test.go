package factory

import (
	"testing"
)

func setup() *Shelf {
	shelf := NewShelf("Hot shelf", "hot", 10)

	return shelf
}

var testItem = Item{
	Id:        "test",
	Name:      "test",
	Temp:      "hot",
	ShelfLife: 10,
	DecayRate: 0.5,
}

func TestRegister(t *testing.T) {
	shelf := setup()
	order := Order{
		Item: testItem,
		Id:   "test-123",
	}

	shelf.Register(order)
	if shelf.FoodOnShelf != 1 {
		t.Errorf("Expected 1, got %d", shelf.FoodOnShelf)
	}

	if !shelf.Contains(order) {
		t.Errorf("Expected true, got false")
	}
}

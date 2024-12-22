package factory

import (
	"math/rand"
)

type Order struct {
	Item Item
}

func (f *Factory) GenerateRandomOrder() Order {
	randomItem := f.Menu.Items[rand.Intn(len(f.Menu.Items))]

	return Order{
		Item: randomItem,
	}
}

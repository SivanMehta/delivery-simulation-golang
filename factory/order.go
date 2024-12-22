package factory

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	Item Item
	Id   string
}

func (f *Factory) GenerateRandomOrder() Order {
	randomItem := f.Menu.Items[rand.Intn(len(f.Menu.Items))]
	now := time.Now().Unix()

	id := fmt.Sprintf(`%s-%d`, randomItem.Id, now)

	return Order{
		Item: randomItem,
		Id:   id,
	}
}

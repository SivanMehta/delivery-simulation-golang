package factory

import (
	"fmt"
	"main/tools"
	"time"
)

type Factory struct {
	Settings tools.SimulationSettings
	Menu     Menu
	Storage  map[string]*Shelf
}

func NewFactory(settings tools.SimulationSettings) Factory {
	menu := GenerateMenu()
	shelves := map[string]*Shelf{
		"hot":      NewShelf("Hot shelf", "hot", 10),
		"cold":     NewShelf("Cold shelf", "cold", 10),
		"frozen":   NewShelf("Frozen shelf", "frozen", 10),
		"overflow": NewShelf("Overflow shelf", "overflow", 15),
	}
	newFactory := Factory{
		Settings: settings,
		Menu:     menu,
		Storage:  shelves,
	}

	return newFactory
}

func (f *Factory) Log(msg string, args ...any) {
	now := time.Now().Unix()
	message := fmt.Sprintf(msg, args...)

	output := fmt.Sprintf(
		`%d,%s,%d,%d,%d,%d`,
		now, message,
		f.Storage["hot"].FoodOnShelf,
		f.Storage["cold"].FoodOnShelf,
		f.Storage["frozen"].FoodOnShelf,
		f.Storage["overflow"].FoodOnShelf,
	)
	fmt.Println(output)
}

func (f *Factory) Intake(order Order) {
	f.Log(`Placing order for %s for %s`, order.Item.Id, order.Item.Temp)

	targetShelf := f.Storage[order.Item.Temp]
	overflowShelf := f.Storage["overflow"]

	if targetShelf.HasCapacity() {
		targetShelf.Register(order)
		f.Log(`Placed order for %s on %s`, order.Item.Id, order.Item.Temp)
	} else if overflowShelf.HasCapacity() {
		overflowShelf.Register(order)
		f.Log(`Placed order for %s on overflow`, order.Item.Id)
	}
}

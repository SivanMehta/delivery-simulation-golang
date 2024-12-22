package factory

import (
	"fmt"
	"main/tools"
)

type Factory struct {
	Settings tools.SimulationSettings
	Menu     Menu
}

func (f *Factory) Intake(order Order) {
	msg := fmt.Sprintf(`Placing order for %s for %s`, order.Item.Id, order.Item.Temp)
	fmt.Println(msg)
}

func NewFactory(settings tools.SimulationSettings) Factory {
	menu := GenerateMenu()
	newFactory := Factory{
		Settings: settings,
		Menu:     menu,
	}

	return newFactory
}

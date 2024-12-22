package factory

import (
	"fmt"
	"main/tools"
)

type Factory struct {
	Settings tools.SimulationSettings
	id       string
}

func (f *Factory) Intake() {
	msg := fmt.Sprintf(`Intaking Order %d at facility #%s`, 123, f.id)
	fmt.Println(msg)
}

func NewFactory(settings tools.SimulationSettings) Factory {
	newFactory := Factory{Settings: settings, id: "ice cream emporium"}

	return newFactory
}

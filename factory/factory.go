package factory

import (
	"fmt"
	"main/tools"
)

type Factory struct {
	Settings tools.SimulationSettings
	Intake   func()
	id       string
}

func Intake() {
	fmt.Println("Intaking order")
}

func NewFactory(settings tools.SimulationSettings) Factory {
	newFactory := Factory{Settings: settings, Intake: Intake, id: "123"}

	return newFactory
}

package main

import (
	"main/factory"
	"main/tools"
	"time"
)

func main() {
	settings := tools.GetSimulationSettings()
	tools.PrintArgs(settings)

	msBetweenOrders := 1000.0 / float32(settings.IngestionRate)
	interval := time.Duration(msBetweenOrders * float32(time.Millisecond))

	facility := factory.NewFactory(settings)

	for {
		factory.GenerateRandomOrder()
		facility.Intake()
		time.Sleep(interval)
	}
}

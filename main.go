package main

import (
	"main/factory"
	"main/tools"
	"time"
)

func main() {
	settings, err := tools.GetSimulationSettings()
	if err != nil {
		panic(err)
	}
	tools.PrintArgs(settings)

	msBetweenOrders := 1000.0 / float32(settings.IngestionRate)
	interval := time.Duration(msBetweenOrders * float32(time.Millisecond))

	facility := factory.NewFactory(settings)

	for {
		order := facility.GenerateRandomOrder()
		facility.Intake(order)
		time.Sleep(interval)
	}
}

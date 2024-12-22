package main

import (
	"fmt"
	"main/tools"
	"time"
)

func main() {
	settings := tools.GetSimulationSettings()

	fmt.Printf(`
Starting simulation with params:
	IngestionRate: %d,
	CourierSpeedLow: %d,
	CourierSpeedHigh: %d

`, settings.IngestionRate, settings.CourierSpeedLow, settings.CourierSpeedHigh)

	msBetweenOrders := 1000.0 / float32(settings.IngestionRate)
	interval := time.Duration(msBetweenOrders * float32(time.Millisecond))

	for {
		fmt.Println("Generating Random Order")
		fmt.Println("Intaking order")
		time.Sleep(interval)
	}
}

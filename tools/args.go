package tools

import (
	"flag"
	"fmt"
)

type SimulationSettings struct {
	IngestionRate    int
	CourierSpeedLow  int
	CourierSpeedHigh int
}

func GetSimulationSettings() SimulationSettings {
	IngestionRate := flag.Int("IngestionRate", 3, "How quickly orders come in, in orders / second")
	CourierSpeedLow := flag.Int("CourierSpeedLow", 2000, "How quickly the fastest couriers can fulfull an order, in ms")
	CourierSpeedHigh := flag.Int("CourierSpeedHigh", 6000, "How slowly couriers can fulfull an order, ms")

	flag.Parse()

	if *CourierSpeedHigh < *CourierSpeedLow {
		msg := fmt.Sprintf(`Invalid courier speed interval: [%d, %d]`, *CourierSpeedLow, *CourierSpeedHigh)
		panic(msg)
	}

	return SimulationSettings{*IngestionRate, *CourierSpeedLow, *CourierSpeedHigh}
}

func PrintArgs(settings SimulationSettings) {
	fmt.Printf(`
Simulation with params:
  IngestionRate: %d
  CourierSpeedLow: %d
  CourierSpeedHigh: %d

time,message,hot,cold,frozen,overflow
`, settings.IngestionRate, settings.CourierSpeedLow, settings.CourierSpeedHigh)
}

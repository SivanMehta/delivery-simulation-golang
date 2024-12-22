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
	CourierSpeedLow := flag.Int("CourierSpeedLow", 2, "How quickly the fastest couriers can fulfull an order, in seconds")
	CourierSpeedHigh := flag.Int("CourierSpeedHigh", 6, "How slowly couriers can fulfull an order, seconds")

	flag.Parse()

	return SimulationSettings{*IngestionRate, *CourierSpeedHigh, *CourierSpeedLow}
}

func PrintArgs(settings SimulationSettings) {
	fmt.Printf(`
###########################
#                         #
# Simulation with params: #
#   IngestionRate: %d      #
#   CourierSpeedLow: %d    #
#   CourierSpeedHigh: %d   #
#                         #
###########################

time,message,hot,cold,frozen,overflow
`, settings.IngestionRate, settings.CourierSpeedLow, settings.CourierSpeedHigh)
}

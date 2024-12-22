package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	ingestionRate := flag.Int("ingestionRate", 3, "How quickly orders come in, in orders / second")
	courierSpeedLow := flag.Int("courierSpeedLow", 2, "How quickly the fastest couriers can fulfull an order, in seconds")
	courierSpeedHigh := flag.Int("courierSpeedHigh", 6, "How slowly couriers can fulfull an order, seconds")

	flag.Parse()

	fmt.Printf(`
Starting simulation with params:
	ingestionRate: %d,
	courierSpeedLow: %d,
	courierSpeedHigh: %d

`, *ingestionRate, *courierSpeedLow, *courierSpeedHigh)

	msBetweenOrders := 1000.0 / float32(*ingestionRate)
	interval := time.Duration(msBetweenOrders * float32(time.Millisecond))

	for {
		fmt.Println("Generating Random Order")
		fmt.Println("Intaking order")
		time.Sleep(interval)
	}
}

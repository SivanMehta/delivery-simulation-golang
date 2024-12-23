package factory

import (
	"fmt"
	"main/tools"
	"maps"
	"math/rand"
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
	acceptedOrder := true

	if targetShelf.HasCapacity() {
		targetShelf.Register(order)
		f.Log(`Placed order for %s on %s`, order.Item.Id, order.Item.Temp)
	} else if overflowShelf.HasCapacity() {
		overflowShelf.Register(order)
		f.Log(`Placed order for %s on overflow`, order.Item.Id)
	} else {
		// attempt to make space
		madeSpace := f.attemptToMakeSpace()
		if madeSpace {
			overflowShelf.Register(order)
			f.Log(`Placed order for %s on overflow`, order.Item.Id)
		} else {
			acceptedOrder = false
			f.Log(`Factory at capacity ¯\_(ツ)_/¯`)
		}
	}

	if acceptedOrder {
		go f.DispatchCourier(order)
	}
}

func (f *Factory) DispatchCourier(order Order) {
	// wait a random amount of time to deliver the order
	// this is basically simulating cooking time
	var low int = f.Settings.CourierSpeedLow
	var high int = f.Settings.CourierSpeedHigh
	deliverySpeed := (low + rand.Intn(high-low))
	interval := time.Duration(deliverySpeed * int(time.Millisecond))
	time.Sleep(interval)

	// now attempt to deliver the order
	targetShelf := f.Storage[order.Item.Temp]
	overflowShelf := f.Storage["overflow"]
	var deliveryStatus string

	if targetShelf.Contains(order) {
		targetShelf.Remove(order)
		deliveryStatus = targetShelf.Name
	} else if overflowShelf.Contains(order) {
		overflowShelf.Remove(order)
		deliveryStatus = overflowShelf.Name
	} else {
		deliveryStatus = "removed"
	}

	if deliveryStatus != "removed" {
		f.Log(`Delivered %s from %s shelf`, order.Item.Id, deliveryStatus)
	} else {
		f.Log("Threw away %s", order.Item.Id)
	}
}

func (f *Factory) attemptToMakeSpace() bool {
	madeSpace := false
	overflow := f.Storage["overflow"]
	for key := range maps.Keys(f.Storage["overflow"].Orders) {
		order := Order{Item: overflow.Orders[key], Id: key}
		if f.Storage[order.Item.Temp].HasCapacity() {
			overflow.Remove(order)
			f.Storage[order.Item.Temp].Register(order)
			f.Log(`Moved %s from overflow to %s`, order.Item.Id, order.Item.Temp)
			madeSpace = true
		}
	}

	return madeSpace
}

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

func (f *Factory) log(msg string, args ...any) {
	var now int64 = time.Now().Unix()
	var message string = fmt.Sprintf(msg, args...)

	var output string = fmt.Sprintf(
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
	f.log(`Placing order for %s for %s`, order.Item.Id, order.Item.Temp)

	var targetShelf *Shelf = f.Storage[order.Item.Temp]
	var overflowShelf *Shelf = f.Storage["overflow"]
	var acceptedOrder bool = true

	if targetShelf.HasCapacity() {
		targetShelf.Register(order)
		f.log(`Placed order for %s on %s`, order.Item.Id, order.Item.Temp)
	} else if overflowShelf.HasCapacity() {
		overflowShelf.Register(order)
		f.log(`Placed order for %s on overflow`, order.Item.Id)
	} else {
		// attempt to make space
		var madeSpace bool = f.attemptToMakeSpace()
		if madeSpace {
			overflowShelf.Register(order)
			f.log(`Placed order for %s on overflow`, order.Item.Id)
		} else {
			acceptedOrder = false
			f.log(`Factory at capacity ¯\_(ツ)_/¯`)
		}
	}

	if acceptedOrder {
		// we do not pass a channel to this goroutine
		// because we do not want to block accepting other orders
		// while accepting this one
		go f.dispatchCourier(order)
	}
}

func (f *Factory) dispatchCourier(order Order) {
	// wait a random amount of time to deliver the order
	// this is basically simulating cooking time
	var low int = f.Settings.CourierSpeedLow
	var high int = f.Settings.CourierSpeedHigh
	var deliverySpeed int = (low + rand.Intn(high-low))
	var interval time.Duration = time.Duration(deliverySpeed * int(time.Millisecond))
	time.Sleep(interval)

	// now attempt to deliver the order
	var targetShelf *Shelf = f.Storage[order.Item.Temp]
	var overflowShelf *Shelf = f.Storage["overflow"]
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
		f.log(`Delivered %s from %s shelf`, order.Item.Id, deliveryStatus)
	} else {
		f.log("Threw away %s", order.Item.Id)
	}
}

func (f *Factory) attemptToMakeSpace() bool {
	var madeSpace bool = false
	var overflow *Shelf = f.Storage["overflow"]
	for key := range maps.Keys(f.Storage["overflow"].Orders) {
		var order Order = Order{Item: overflow.Orders[key], Id: key}
		if f.Storage[order.Item.Temp].HasCapacity() {
			overflow.Remove(order)
			f.Storage[order.Item.Temp].Register(order)
			f.log(`Moved %s from overflow to %s`, order.Item.Id, order.Item.Temp)
			madeSpace = true
		}
	}

	return madeSpace
}

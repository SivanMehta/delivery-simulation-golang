# Challenge Prompt
Create a real-time system that emulates the fulfillment of delivery orders for a
kitchen. The system should implement a real-time simulation as opposed to a
discrete-event simulation. The kitchen should receive 2 delivery orders per
second (see Appendix on Orders). The kitchen should instantly cook the order upon
receiving it, and then place the order on the best-available shelf
(see Appendix on Shelves) to await pick up by a courier.

Upon receiving an order, the system should dispatch a courier to pick up and
deliver that specific order. The courier should arrive randomly between 2-6
seconds later. The courier should instantly pick up the order upon arrival. Once
an order is picked up, the courier should instantly deliver it. You can use any
programming language, framework, and IDE to demonstrate your best work; however,
we strongly discourage the use of microservices, kafka, REST APIs, RPCs, DBs,
etc due to time constraints.

## Appendix

### Orders

Orders must be parsed from the file and ingested into your system at a rate of 2
orders per second. You are expected to make your order ingestion rate configurable,
so that we can test your system’s behavior with different ingestion rates.

### Shelves

The kitchen pick-up area has multiple shelves to hold cooked orders at different
temperatures. Each order should be placed on a shelf that matches the order’s
temperature. If that shelf is full, an order can be placed on the overflow shelf.
If the overflow shelf is full, an existing order of your choosing on the
overflow should be moved to an allowable shelf with room. If no such move is
possible, a random order from the overflow shelf should be discarded as waste
(and will not be available for a courier pickup). The following table details
the kitchen’s shelves:

| Name | Allowable Temperature(s) | Capacity
| -- | -- | -- |
|Hot shelf | Hot | 10
|Cold shelf | Cold | 10
|Frozen shelf | Frozen | 10
|Overflow shelf | Any temperature | 15

## Shelf Life (extra credit)

Instead of orders persisting indefinitely on a shelf, now consider that orders
have an inherent value that will deteriorate over time, based on the order’s
shelfLife and decayRate fields. Orders that have reached a value of zero are
considered wasted: they should never be delivered and should be removed from
the shelf. Please display the current order value when displaying an order in
your system’s output.

```
value = (shelfLife - orderAge - (decayRate * orderAge * shelfDecayModifier))
       _____________________________________________________________________
                              shelfLife
```

`shelfDecayModifier` is `1` for single-temperature shelves and `2` for the overflow shelf.

##  Build Instructions

> This assumes you already have [Golang](https://go.dev/learn/) installed.

```shell
# installs dependencies
go mod download

# runs the tests and generates a coverage report
go test

# build the simulation
go build

# start the simulation
./main
```

You can tune the simulation by providing CLI parameters. Here are all of the
flags with their default value:

```shell
./main
  # how quickly orders come in, in orders / second
  -IngestionRate=3
  
  # the fastest couriers can fulfull an order, in seconds
  -CourierSpeedLow=2
  
  # the slowest couriers can fulfull an order, seconds
  -CourierSpeedHigh=6
```

package factory

import (
	"fmt"
)

type Order struct {
}

func GenerateRandomOrder() Order {
	fmt.Println("generating order from JSON")
	return Order{}
}

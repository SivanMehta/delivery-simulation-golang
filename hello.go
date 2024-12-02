package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	var msg string = quote.Go()
	fmt.Println(msg)
}

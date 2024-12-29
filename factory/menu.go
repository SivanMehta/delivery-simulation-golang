package factory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Item struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Temp      string  `json:"temp"`
	ShelfLife int     `json:"shelfLife"`
	DecayRate float32 `json:"decayRate"`
}

type Menu struct {
	Items []Item `json:"items"`
}

func GenerateMenu() Menu {
	file, err := os.Open("factory/menu.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our file so that we can parse it first
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var menu Menu
	json.Unmarshal(byteValue, &menu)

	return menu
}

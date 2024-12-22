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
	jsonFile, err := os.Open("factory/menu.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened menu.json")
	}
	// defer the closing of our jsonFile so that we can parse it first
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var menu Menu
	json.Unmarshal(byteValue, &menu)

	return menu
}

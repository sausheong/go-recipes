package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Recipe struct {
	Id          int
	Name        string
	Ingredients []string
}

// write data
func write(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

// read the data
func read(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	recipe := Recipe{
		Id:          1,
		Name:        "Chocolate Tart",
		Ingredients: []string{"Flour", "Sugar", "Eggs", "Butter", "Chocolate"},
	}
	write(recipe, "chocolate_tart")
	var r Recipe
	read(&r, "chocolate_tart")
	fmt.Println(r)
}

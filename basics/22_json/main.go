package main

import "encoding/json"

type Person struct {
	name  string
	addr  string
	phone int
}

func main() {
	// JSON Marshalling -> returns JSON representation as []byte
	p1 := Person{name: "John", addr: "Nepal", phone: 1234}
	barr, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	var p2 Person
	// JSON Unmarshal -> converts a JSON []byte into a Go object
	err = json.Unmarshal(barr, &p2)
}

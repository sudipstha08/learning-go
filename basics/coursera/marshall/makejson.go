package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Address string	
}

func main() {
	fmt.Println("Please enter your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Please enter your address")

	var address string
	fmt.Scanln(&address)

	p1 := &Person{name, address} 
	person, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(person))
}

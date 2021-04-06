// Write a program which prompts the user to first enter a name, 
// and then enter an address. Your program should create a map and 
// add the name and address to the map using the keys “name” and “address”, 
// respectively. Your program should use Marshal() to create a JSON object
// from the map, and then your program should print the JSON object.

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

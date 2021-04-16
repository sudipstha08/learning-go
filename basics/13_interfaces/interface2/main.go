package main

import "fmt"

type Animal interface{
	// The speak methd takes no argument & returns a string
	Speak() string 
}

type Dog struct {}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {}

func (c Cat) Speak() string {
	return "Meow!"
}

type Horse struct {}

func (c Horse) Speak() string {
	return "Neigh!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Horse{}}
	for _, animal := range animals {
		fmt.Println("animal",animal)
		fmt.Println("sound:", animal.Speak())
	}
}
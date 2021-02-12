package main

import (
	"fmt"
	"learning-go/03_packages/grettings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
package main

import (
	"fmt"
	"learning-go/basics/03_packages/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}

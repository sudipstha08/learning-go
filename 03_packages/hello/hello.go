package main

import (
	"fmt"
	"go_crash_course/03_packages/grettings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
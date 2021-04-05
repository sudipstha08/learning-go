package main

import "fmt"

// Array literal

// An array predefined with values
// length of literal must be the length of the array

func main() {
	x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment

	fmt.Println(x)
	fmt.Println(y)
}


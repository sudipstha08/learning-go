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

	// [...] in array literals infers size from number of initializers
	z := [...]int{1, 2, 3, 4, 5}    // size of z is 5
	fmt.Println(z)

	for idx, value := range z {
		fmt.Printf("index: %d , value: %d \n", idx, value)
	}
}

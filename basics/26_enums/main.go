package main

import "fmt"

// iota is an identifier that is used with constant and which can 
// simplify constant definitions that use auto-increment numbers
const (
	a0 = iota
	a1 = iota
	a2 = iota
)

const (
	b0 = iota
	b1
	b2
)

// To start a list of constants at 1 instead of 0
const (
	c0 = iota + 1
	c1
	c2
)

// you can use the blank identifier to skip a value in a list of constants.
const (
	d1 = iota + 1
	_
	d3
	d4
)

func main() {
	fmt.Println(a0, a1, a2) //Print : 0 1 2
	fmt.Println(b0, b1, b2) //Print : 0 1 2
	fmt.Println(c0, c1, c2) //Print : 1 2 3
	fmt.Println(d1, d3, d4) //Print : 1 3 4
}
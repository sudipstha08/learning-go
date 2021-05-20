package main

import "fmt"

type T struct {
	name string
}

//  Method is a function with additional single-element list
// of parameters called receiver. It’s placed right before
// the name of a method.

// You can only define methods on a type defined in that same package. 
// Here T must be on the same package as PrintName method
func (t T) PrintName() {
	fmt.Println(t.name)
}
func main() {
	t := T{name: "Michał"}
	t.PrintName()
}

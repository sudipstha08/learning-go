package main

import "fmt"

type T struct {
	name string
}

//  Method is a function with additional single-element list
// of parameters called receiver. It’s placed right before
// the name of a method.

func (t T) PrintName() {
	fmt.Println(t.name)
}
func main() {
	t := T{name: "Michał"}
	t.PrintName()
}

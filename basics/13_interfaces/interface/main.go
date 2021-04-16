// An interface is two things: it is a set of methods, but it is also a type
package main

import "fmt"

type Envelope struct {
	Body `xml:"soap:"`
}

// The interface{} type, the empty interface is the interface that has no methods.
// interface{} means you can put value of any type, including your own custom type. All types in Go satisfy an empty interface (interface{} is an empty interface).
type Body struct {
	Msg interface{}
}

func main() {
	b := Body{}
	b.Msg = "5"
	fmt.Printf("%#v %T \n", b.Msg, b.Msg) // Output: "5" string
	b.Msg = 5

	fmt.Printf("%#v %T", b.Msg, b.Msg) //Output:  5 int
}

// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

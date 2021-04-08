// You are allowed to Create a method with non-struct type receiver as long 
// as the type and the method definitions are present in the same package.
package main

import "fmt"

// Type defination
type data int

// NON STRUCT RECEIVER TYPE
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

func main() {
	value1 := data(23)
	value2 := data(22)
	result := value1.multiply(value2)
	fmt.Println("Result: ", result)
}
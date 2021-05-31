package main

import "fmt"

/* GLOBAL VARIABLE DECLARATION */
var g int = 10

func main() {
	/* LOCAL VARIABLE DECLARATION */
	var a, b, c int

	/* INITIALIZATION */
	a = 10
	b = 20
	c = a + b

	fmt.Printf("value of a = %d, b = %d and c = %d\n and g = %d", a, b, c, g)

	// Formal parameters are treated as local variables with-in that function and 
	// they take preference over the global variables
	var g int = 10
	fmt.Printf("value of a in main() = %d\n",  g);
}

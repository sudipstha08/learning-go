package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Brad"
	var age int32 = 37
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	// name := "Brad"
	// email := "brad@gmail.com"

	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)

	// Multiple variables can be declared in one statement.
	// var var_1, var_2 = value_1, value_2

	//	you can use the multiple-line variable declaration syntax.
	// const (
		// const_1 = value_1
		// const_2 = value_2
	// )


	// iota that can be used when declaring enumerated constants. 
	// This keyword yields an incremented value by 1 starting from 0, 
	// each time it is used.
	const (
    a = iota // a == 0
    b = iota // b == 1
    c = iota // c == 2
    d        // d == 3 (implicitely d = iota)
)
}

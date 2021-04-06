package main

import "fmt"

func main() {
	// Slice is an window on an underlying array
	// Pointer indicates the start of the slice
	// length is the number of elements in the slice
	// Capacity is the max number of elements
  // Size of slice can be increased upto the size of the array
	// The first index position in a slice is always 0 and 
	// the last one will be (length of slice – 1).

	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3] // ["b", "c"]
	s2 := arr[2:5] // ["c", "d", "e"]
	fmt.Println(s1)
	fmt.Println(s2)
}
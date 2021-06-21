package main

import "fmt"

func main() {
	// Size of a slice can be increased by append
	sli := make([]int, 0, 3)
	fmt.Println("length", len(sli))
	fmt.Println("capacity", cap(sli))

	// when the new elements append to the slice a new array is created
	sli = append(sli, 100)
	sli = append(sli, 130)
	sli = append(sli, 140)
	sli = append(sli, 160)
	fmt.Println(sli)

	var s []int
	s = append(s, 0)
	s = append(s, 1)
	fmt.Println(s)
}

package main

import "fmt"

func main() {
	var i int16 = 223
	var j int8
	j = int8(i)
	fmt.Printf("i = %d and j = %d", i, j)
	fmt.Println("")
}

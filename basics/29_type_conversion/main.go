// TYPE Conversion || TYPE casting
package main

import "fmt"

func main() {
	var i int16 = 223
	var j int8
	j = int8(i)
	fmt.Printf("i = %d and j = %d", i, j)
	fmt.Println("")

	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("Value of mean : %f\n", mean)
}

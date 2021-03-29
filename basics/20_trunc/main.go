package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter the floating point number")

	var floatNum float64

	fmt.Scanln(&floatNum)
	decimalNum := math.Trunc(floatNum)
	fmt.Printf("The truncated number is %g", decimalNum)
}

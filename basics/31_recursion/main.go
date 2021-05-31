package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {
	fmt.Println("Please enter a number")
	var inputNum int
	fmt.Scanln(&inputNum)
	fmt.Println("-----Calculated by recursive method------")
	fmt.Printf("Factorial of %d is %d\n", inputNum, factorial(inputNum))
}

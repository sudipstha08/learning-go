package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please enter the string")
	var inputString string
	fmt.Scanln(&inputString)
	startsWith := strings.HasPrefix(strings.ToUpper(inputString), "I")
	endsWith := strings.HasSuffix(strings.ToUpper(inputString), "N")
	contains := strings.Contains(strings.ToUpper(inputString), "A")

	if startsWith && endsWith && contains {
		println("Found!")
	} else {
		println("Not Found!")
	}
}

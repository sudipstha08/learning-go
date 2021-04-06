// Write a program which prompts the user to enter a string. The program searches 
// through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program 
// should print “Found!” if the entered string starts with the character ‘i’, 
// ends with the character ‘n’, and contains the character ‘a’. The program should 
// print “Not Found!” otherwise. The program should not be case-sensitive, so it
// does not matter if the characters are upper-case or lower-case.

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

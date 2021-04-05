// The panic keyword is used to terminate the program with a
// custom error message. When the panic keyword is encountered,
// the following set of instructions are followed:
// 1. Execution for the current function stops
// 2. Any function called with defer keyword is executed
// 3. The program execution is terminated

package main

import "fmt"


func recoveryFunction() {
  fmt.Println("This is recovery function...")
}

// whenever a panic function is introduced, it executes all 
// the defer functions associated with the currently executing thread. 
// The defer function can be used to free up the resources. defer functions 
// are executed just before the termination of the currently executing function
func executePanic() {
	defer recoveryFunction()
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}

func main() {
	executePanic()
	fmt.Println("Main block is executed completely...")
}

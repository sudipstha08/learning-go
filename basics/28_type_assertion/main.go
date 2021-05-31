// Assertion means a statement that you strongly believe is true.
// Type assertion works only with interface
// i.(TYPE) only works with switch case
package main

import "fmt"

func main() {
	var i interface{} = 5
	val, ok := i.(int)
	fmt.Println("value", val)
	fmt.Println("ok", ok)

	// switch i.(TYPE)
  // case: int
  //   fmt.Printf("This type belongs to integer")
  //   break;
  // case: int32
  //   fmt.Printf("This type belongs to integer 32 bits")
  //   break;
  // case: int64
  //   fmt.Printf("This type belongs to integer 64 bits")
  //   break;
}

// With the help of a pointer receiver if a change made in the method
// will reflect in the caller which is not possible with the value receiver.
package main

import "fmt"

type Author struct {
	Name     string
	Branch   string
	Articles int
}

// METHOD WITH A POINTER RECEIVER
func (a *Author) show(abranch string) {
	fmt.Println("Params", abranch)
	fmt.Println("Receiver", a)
	(*a).Branch = abranch
}

func main() {
	author1 := Author{Name: "William", Branch: "Math"}

	fmt.Println("Author's name: ", author1.Name)
	fmt.Println("Branch: ", author1.Branch)
	// Creating a pointer
	a := &author1

	// Calling the show method
	a.show("Science")
	fmt.Println("Author's name: ", author1.Name)
	fmt.Println("Branch: ", author1.Branch)
}

// when a function has a value argument, then it will only accept the
// values of the parameter, and if you try to pass a pointer to a value
// function, then it will not accept or vice versa. But a Go method can
// accept both value and pointer whether it is defined with pointer or value receiver.

// Go program to illustrate how the
// method can accept pointer and value

package main

import "fmt"

// Author structure
type Author struct {
	name   string
	branch string
}

// Method with a pointer receiver of author type
func (a *Author) show_1(abranch string) {
	(*a).branch = abranch
}

// Method with a value receiver of author type
func (a Author) show_2() {
	a.name = "Gourav"
	fmt.Println("Author's name(Before) : ", a.name)
}

func main() {
	// Initializing the values of the author structure
	res := Author{
		name:   "Sona",
		branch: "CSE",
	}

	fmt.Println("Branch Name(Before): ", res.branch)

	// Calling the show_1 method (pointer method) with value
	res.show_1("ECE")
	fmt.Println("Branch Name(After): ", res.branch)

	// Calling the show_2 method (value method) with a pointer
	(&res).show_2()
	fmt.Println("Author's name(After): ", res.name)
}

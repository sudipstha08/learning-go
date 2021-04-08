//Methods With Same Name ->
// it is allowed to create two or more methods with the same name
// in the same package, but the receiver of these methods must be of different types
package main

import "fmt"

// Creating structures
type student struct {
	name   string
	branch string
}

type teacher struct {
	language string
	marks    int
}

// Same name methods, but with different type of receivers
func (s student) show() {
	fmt.Println("Name of the Student:", s.name)
	fmt.Println("Branch: ", s.branch)
}

func (t teacher) show() {
	fmt.Println("Language:", t.language)
	fmt.Println("Student Marks: ", t.marks)
}

func main() {

	val1 := student{"Rohit", "EEE"}

	val2 := teacher{"Java", 50}

	val1.show()
	val2.show()
}

package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// Modify the object itself
func (s *Student) setAge(age int) {
	s.age = age
}

// Changes only the local copy
func (s Student) getAge(age int) {
	s.age = age
	fmt.Println("get local age:", s.age)
}

func main() {
	s1 := Student{"Tim", []int{60, 40, 90}, 20}
	fmt.Println(s1)
	s1.getAge(7)
	fmt.Println(s1)

	fmt.Println(s1)
	s1.setAge(7)
	fmt.Println(s1)
}

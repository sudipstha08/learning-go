package main

type Person struct {
	FirstName string
	LastName string
	Number int32
	Address string
	email string
}

// Constructor
// This func returns a pointer to repo
func New() *Person {
	return &Person{}
}



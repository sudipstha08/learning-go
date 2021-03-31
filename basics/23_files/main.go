package main

import (
	"fmt"
	"io/ioutil"
)

// BASIC OPERATIONS
// Open -> get handle for access
// Read -> read bytes into []byte
// Write -> write []byte into file
// Close -> release handle
// Seek -> move read/write head

func main() {
	// data is []byte filled with contents of the entire file
	// Explicit open and close is not needed
	data, error := ioutil.ReadFile("test.txt")
	if error == nil {
		fmt.Println(data)
	}

	// Writes []byte to file
	// Creates a file
	// Unix styles permission bytes
	const dat = "This is a WriteFile test"
	err := ioutil.WriteFile("output-file.txt", data, 0777)
	if err == nil {
		fmt.Println("File successfully written !")
	}
}

// OS package file access
package main

import (
	"fmt"
	"os"
)

func main() {
	//opens a file and returns a file descriptor
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	
	barr := make([]byte, 10)
	// Closes a file
	// Reads from a file into a []byte
	nb, err := file.Read(barr)
	fmt.Println(nb)

	// Writes a string
	nb, err = file.WriteString("Hello Hello")
	file.Close()

	// write a []byte into a file
	// os.Write()
}

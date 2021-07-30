// A defer statement defers the execution of a function until the
// surrounding function returns.

// The deferred call's arguments are evaluated immediately, but the 
// function call is not executed until the surrounding function returns.

// Usually for purpose of cleanup e.g. ensure and finally 
// would be used in other languages.

package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// It’s important to check for errors when closing a file, even in a deferred function.

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

// https://medium.com/a-journey-with-go/go-how-does-defer-statement-work-1a9492689b6e
// https://www.youtube.com/watch?v=t4Ps0yXEdXM
// https://www.geeksforgeeks.org/defer-keyword-in-golang/#:~:text=In%20Go%20language%2C%20defer%20statements,until%20the%20nearby%20functions%20returns.
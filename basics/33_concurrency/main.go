// Concurrency in Golang is the ability for functions to run independent
// of each other. A goroutine is a function that is capable of running
// concurrently with other functions. When you create a function as a goroutine,
// it has been treated as an independent unit of work that gets scheduled and
// then executed on an available logical processor

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

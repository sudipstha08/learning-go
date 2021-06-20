// CONCURRENCY WITH PARALLELSIM
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// ADDING AN EXECUTION CORE
	runtime.GOMAXPROCS(4)
	start := time.Now()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("FIRST: ", i)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("SECOND: ", i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)
}

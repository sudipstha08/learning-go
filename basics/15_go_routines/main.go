// GO ROUTINES - CONCURRENCY && PARALLELISM

package main

import (
	"fmt"
	"time"
)

// ADD GOROUTINE BY USING KEYWORD "go"
// CONCURRENCY
func main() {
	start := time.Now()
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)
}

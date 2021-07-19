package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	formatTime := t.Format("20060102150405")
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(formatted)
	now := time.Now().String()
	fmt.Println(formatTime)
	fmt.Println("time now", now)
	// elapsed := now.Sub(t)
	// fmt.Println("elapsed", elapsed)
}

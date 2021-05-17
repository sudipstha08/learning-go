package main

import "fmt"

func PrintMap(c map[string]string) {
	// iterating over maps
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func main() {
	// var colors1 map[string]string // OR
	colors := make(map[string]string)
	colors["white"] = "#fff"
	colors["black"] = "#00000"
	colors["red"] = "#110000"
	m := len(colors)
	PrintMap(colors)
	fmt.Println(m)
}

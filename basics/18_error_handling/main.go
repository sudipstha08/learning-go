// Golang provides a simple interface for error
package main

import (
	"errors"
)

type error interface {
	Error() string
}

func calculateArea(radius int) (int, error) {
	if radius < 0 {
		return 0, errors.New("Provide Positive Number")
	}
	return radius * radius, nil
}

func main() {
	calculateArea(5)
}

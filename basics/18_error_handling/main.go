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
		return nil, errors.New("Provide Positive Number")
	}
	return radius * radius, nil
}

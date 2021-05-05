package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type widget struct {
	id string
}

type Widget interface {
	// ID returns a widget's unique identifier
	ID() string
}

func (w widget) ID() string {
	return w.id
}

//  “constructor-like” function to handle initialization:
func NewWidget() Widget {
	return widget{
		id: uuid.NewV4().String(),
	}
}

func main() {
	w := NewWidget()
	fmt.Println("id : ", w.ID())
}

// Names that are capitalized are exported ("public," essentially) 
// Names that are not capitalized are unexported ("package private,") - not accessible to othre packages	
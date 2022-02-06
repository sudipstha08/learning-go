package widgets

import uuid "github.com/satori/go.uuid"

type widget struct {
	id string
}

type Widget interface {
	ID() string
}

// The constructor-like NewWidget() function returns an instance of an unexported type.
func NewWidget() widget {
	return widget{
		id: uuid.NewV4().String(),
	}
}

func (w widget) ID() string {
	return w.id
}

// https://krancour.medium.com/go-pointers-why-i-use-interfaces-in-go-338ae0bdc9e4

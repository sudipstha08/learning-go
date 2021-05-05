package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func getGeometry(s geometry) float64 {
	return s.area()
}

func main() {
	r := rect{3, 4}
	c := circle{radius: 5}
	measure(&r)
	fmt.Println("----------Circle---------")
	measure(c)

	geometries := []geometry{&r, c}
	for _,geometry := range geometries{
		// fmt.Println("area", geometry.area())
		// fmt.Println("perim", geometry.perim())
		fmt.Println(getGeometry(geometry))
	} 
}

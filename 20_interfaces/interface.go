package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type surface interface {
	area() float64
}

func printArea(s surface) {
	fmt.Println("Area is:", s.area())
}

func main() {
	// In Go, interface implementation is implicit. That means that a structed type doesn't
	// implement an interface directly like traditional OO languages, but if for example a
	// function demands an interface type in its arguments and the parameter type has the same
	// signature method, than it is assumed that the parameter implements the required interface.

	r1 := rectangle{10, 20}
	printArea(r1)

	c1 := circle{20}
	printArea(c1)
}

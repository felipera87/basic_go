package surfaces

import (
	"math"
)

// Rectangle represents the squared surface
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a squared surface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents the circumference surface
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Surface is an interface that defines a surface of any type
type Surface interface {
	area() float64
}

package structs

import "math"

// Shape represents any geometric object that has an area.
type Shape interface {
	Area() float64
}

// Rectangle represents a geometric object defined solely by its width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a geometric object defined solely by its radius.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

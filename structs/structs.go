package structs

import "math"

// Shape specifies the accepted shapes
type Shape interface {
	Area() float64
}

// Rectangle object
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates with the width and height of floats64
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle object
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle object
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

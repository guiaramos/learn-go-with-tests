package maths

import (
	"math"
	"time"
)

// Point represents a two dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	clockCenterX     = 150
	clockCenterY     = 150
)

// SecondHand is the unit vector of the second hand of an analogue clock at time `t` represented as a Pointer.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{X: p.X * secondHandLength, Y: p.Y * secondHandLength}
	p = Point{X: p.X, Y: -p.Y}
	p = Point{X: p.X + clockCenterX, Y: p.Y + clockCenterY}
	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

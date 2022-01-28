package clockface

import (
	"math"
	"time"
)

// A Point represents a 2D Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue
// clock represented as a Point.
func SecondHand(t time.Time) Point {
	length := float64(90)
	xReference := float64(150)
	yReference := float64(150)

	p := secondHandPoint(t)
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + xReference, p.Y + yReference}

	return p
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / (float64(t.Second())))
}

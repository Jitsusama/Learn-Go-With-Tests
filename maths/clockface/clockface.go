package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

// A Point represents a 2D Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

func SvgWriter(w io.Writer, t time.Time) {
	secondHand := secondHand(t)
	io.WriteString(w, fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
	<svg xmlns="http://www.w3.org/2000/svg"
		 width="100%%"
		 height="100%%"
		 viewBox="0 0 300 300"
		 version="2.0">
		<circle cx="150" cy="150" r="100"
			  style="fill:#fff;stroke:#000;stroke-width:5px;"/>
		<line x1="150" y1="150" x2="%.3f" y2="%.3f"
			  style="fill:none;stroke:#f00;stroke-width:3px;"/>
	</svg>
	`, secondHand.X, secondHand.Y))
}

func secondHand(t time.Time) Point {
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

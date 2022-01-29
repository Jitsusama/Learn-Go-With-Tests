package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

// A Vector represents a 2D Cartesian coordinate.
type Vector struct {
	X float64
	Y float64
}

func SvgWriter(w io.Writer, t time.Time) {
	hourHand := hourHand(t)
	minuteHand := minuteHand(t)
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
    <line x1="150" y1="150" x2="%.3f" y2="%.3f" title="seconds"
          style="fill:none;stroke:#f00;stroke-width:2px;"/>
    <line x1="150" y1="150" x2="%.3f" y2="%.3f" title="minutes"
          style="fill:none;stroke:#000;stroke-width:3px;"/>
    <line x1="150" y1="150" x2="%.3f" y2="%.3f" title="hours"
          style="fill:none;stroke:#000;stroke-width:4px;"/>
</svg>`, secondHand.X, secondHand.Y, minuteHand.X, minuteHand.Y, hourHand.X, hourHand.Y))
}

func hourHand(t time.Time) Vector {
	length := float64(50)
	xReference := float64(150)
	yReference := float64(150)

	p := hourHandVector(t)
	p = Vector{p.X * length, p.Y * length}
	p = Vector{p.X, -p.Y}
	p = Vector{p.X + xReference, p.Y + yReference}

	return p
}

func hourHandVector(t time.Time) Vector {
	return convertAngle(hoursInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour()%12)))
}

func minuteHand(t time.Time) Vector {
	length := float64(80)
	xReference := float64(150)
	yReference := float64(150)

	p := minuteHandVector(t)
	p = Vector{p.X * length, p.Y * length}
	p = Vector{p.X, -p.Y}
	p = Vector{p.X + xReference, p.Y + yReference}

	return p
}

func minuteHandVector(t time.Time) Vector {
	return convertAngle(minutesInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func secondHand(t time.Time) Vector {
	length := float64(90)
	xReference := float64(150)
	yReference := float64(150)

	p := secondHandVector(t)
	p = Vector{p.X * length, p.Y * length}
	p = Vector{p.X, -p.Y}
	p = Vector{p.X + xReference, p.Y + yReference}

	return p
}

func secondHandVector(t time.Time) Vector {
	return convertAngle(secondsInRadians(t))
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / (float64(t.Second())))
}

func convertAngle(angle float64) Vector {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Vector{x, y}
}

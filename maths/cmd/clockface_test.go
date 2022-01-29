package main

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/jitsusama/lgwt/maths/clockface"
)

func TestSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			parseTime("2022-01-29T00:00:00Z"),
			Line{X1: 150, X2: 150, Y1: 150, Y2: 60},
		},
		{
			parseTime("2022-01-29T00:00:30Z"),
			Line{X1: 150, X2: 150, Y1: 150, Y2: 150 + 90},
		},
	}
	for _, c := range cases {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			buffer := bytes.Buffer{}
			svg := Svg{}

			clockface.SvgWriter(&buffer, c.time)
			xml.Unmarshal(buffer.Bytes(), &svg)

			if len(svg.Line) < 1 || svg.Line[0] != c.line {
				t.Errorf("want %v got %v", c.line, svg.Line)
			}
		})
	}
}

func TestMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			parseTime("2022-01-29T00:30:00Z"),
			Line{X1: 150, X2: 150, Y1: 150, Y2: 230},
		},
		{
			parseTime("2022-01-29T00:45:00Z"),
			Line{X1: 150, X2: 70, Y1: 150, Y2: 150},
		},
	}
	for _, c := range cases {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			buffer := bytes.Buffer{}
			svg := Svg{}

			clockface.SvgWriter(&buffer, c.time)
			xml.Unmarshal(buffer.Bytes(), &svg)

			if len(svg.Line) < 2 || svg.Line[1] != c.line {
				t.Errorf("want %v got %v", c.line, svg.Line)
			}
		})
	}
}

func parseTime(isoTime string) time.Time {
	ts, _ := time.Parse(time.RFC3339, isoTime)
	return ts
}

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

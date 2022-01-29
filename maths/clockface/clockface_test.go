package clockface

import (
	"math"
	"testing"
	"time"
)

func TestHourHandVector(t *testing.T) {
	cases := []struct {
		time   time.Time
		vector Vector
	}{
		{parseTime("06:00:00"), Vector{0, -1}},
		{parseTime("21:00:00"), Vector{-1, 0}},
	}
	for _, c := range cases {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			if g := hourHandVector(c.time); !vectorsEqual(g, c.vector) {
				t.Fatalf("want %v vector got %v vector", c.vector, g)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{parseTime("06:00:00"), math.Pi},
		{parseTime("00:00:00"), 0},
		{parseTime("21:00:00"), math.Pi * 1.5},
		{parseTime("00:01:30"), math.Pi / ((6 * 60 * 60) / 90)},
	}
	for _, c := range cases {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			if g := hoursInRadians(c.time); !radiansEqual(g, c.angle) {
				t.Fatalf("want %v radians got %v radians", c.angle, g)
			}
		})
	}
}

func TestMinuteHandVector(t *testing.T) {
	conversions := []struct {
		time   time.Time
		vector Vector
	}{
		{parseTime("00:30:00"), Vector{X: 0, Y: -1}},
		{parseTime("00:45:00"), Vector{X: -1, Y: 0}},
	}
	for _, c := range conversions {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			if g := minuteHandVector(c.time); !vectorsEqual(g, c.vector) {
				t.Fatalf("want %v point got %v point", c.vector, g)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	conversions := []struct {
		time  time.Time
		angle float64
	}{
		{parseTime("00:30:00"), math.Pi},
		{parseTime("00:00:07"), 7 * (math.Pi / (30 * 60))},
	}
	for _, c := range conversions {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			if g := minutesInRadians(c.time); !radiansEqual(g, c.angle) {
				t.Fatalf("want %v radians got %v radians", c.angle, g)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	conversions := []struct {
		time   time.Time
		vector Vector
	}{
		{parseTime("00:00:30"), Vector{X: 0, Y: -1}},
		{parseTime("00:00:45"), Vector{X: -1, Y: 0}},
	}
	for _, c := range conversions {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			if g := secondHandVector(c.time); !vectorsEqual(g, c.vector) {
				t.Fatalf("want %v point got %v point", c.vector, g)
			}
		})
	}
}

func TestSecondsInRadians(t *testing.T) {
	conversions := []struct {
		time  time.Time
		angle float64
	}{
		{parseTime("00:00:30"), math.Pi},
		{parseTime("00:00:00"), 0},
		{parseTime("00:00:45"), (math.Pi / 2) * 3},
		{parseTime("00:00:07"), (math.Pi / 30) * 7},
	}
	for _, c := range conversions {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			if g := secondsInRadians(c.time); !radiansEqual(g, c.angle) {
				t.Fatalf("want %v radians got %v radians", c.angle, g)
			}
		})
	}
}

func vectorsEqual(a Vector, b Vector) bool {
	return math.Abs(a.X-b.X) < 1e-7 && math.Abs(a.Y-b.Y) < 1e-7
}

func radiansEqual(a float64, b float64) bool {
	return math.Abs(a-b) < 1e-7
}

func parseTime(timestamp string) time.Time {
	value, _ := time.Parse("15:04:05", timestamp)
	return value
}

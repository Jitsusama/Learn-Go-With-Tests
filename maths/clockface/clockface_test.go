package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	conversions := []struct {
		time  time.Time
		angle float64
	}{
		{timestamp("00:00:30"), math.Pi},
		{timestamp("00:00:00"), 0},
		{timestamp("00:00:45"), (math.Pi / 2) * 3},
		{timestamp("00:00:07"), (math.Pi / 30) * 7},
	}
	for _, c := range conversions {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			if g := secondsInRadians(c.time); g != c.angle {
				t.Fatalf("want %v radians got %v radians", c.angle, g)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	conversions := []struct {
		time  time.Time
		point Point
	}{
		{timestamp("00:00:30"), Point{1.2246467991473515e-16, -1}},
		{timestamp("00:00:45"), Point{-1, -1.8369701987210272e-16}},
	}
	for _, c := range conversions {
		t.Run(c.time.Format("15:04:05"), func(t *testing.T) {
			if g := secondHandPoint(c.time); g != c.point {
				t.Fatalf("want %v point got %v point", c.point, g)
			}
		})
	}
}

func timestamp(timestamp string) time.Time {
	value, _ := time.Parse("15:04:05", timestamp)
	return value
}

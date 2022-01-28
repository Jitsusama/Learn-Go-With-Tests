package acceptance

import (
	"testing"
	"time"

	"github.com/jitsusama/lgwt/maths/clockface"
)

func TestSecondHandAtMidnight(t *testing.T) {
	ts, _ := time.Parse(time.RFC3339, "2022-01-27T00:00:00Z")

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(ts)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	ts, _ := time.Parse(time.RFC3339, "2022-01-27T00:00:30Z")

	want := clockface.Point{X: 150, Y: 150 + 90}
	got := clockface.SecondHand(ts)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

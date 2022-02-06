package testing

import (
	"io"
	"testing"
	"time"
)

func NewSpiedGame(blindAlert []byte) *SpiedGame {
	return &SpiedGame{blindAlert: blindAlert}
}

type SpiedGame struct {
	started      bool
	startedWith  int
	finishedWith string
	blindAlert   []byte
}

func (g *SpiedGame) Start(players int, alerter io.Writer) {
	g.started = true
	g.startedWith = players
	alerter.Write(g.blindAlert)
}

func (g *SpiedGame) Finish(winner string) {
	g.finishedWith = winner
}

func (g *SpiedGame) AssertStart(t *testing.T, players int) {
	t.Helper()
	if !retry(500*time.Millisecond, func() bool {
		return g.startedWith == players
	}) {
		t.Errorf("start: got %d want %d", g.startedWith, players)
	}
}

func (g *SpiedGame) AssertFinish(t *testing.T, winner string) {
	t.Helper()
	if !retry(500*time.Millisecond, func() bool {
		return g.finishedWith == winner
	}) {
		t.Errorf("finish: got %q want %q", g.finishedWith, winner)
	}
}

func retry(d time.Duration, do func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if do() {
			return true
		}
	}
	return false
}

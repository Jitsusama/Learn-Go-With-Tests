package testing

import (
	"io"
	"testing"
)

type SpiedGame struct {
	started      bool
	startedWith  int
	finishedWith string
}

func (g *SpiedGame) Start(players int, alertDestination io.Writer) {
	g.started = true
	g.startedWith = players
}

func (g *SpiedGame) Finish(winner string) {
	g.finishedWith = winner
}

func (g *SpiedGame) AssertStart(t *testing.T, players int) {
	t.Helper()
	if g.startedWith != players {
		t.Errorf("start: got %d want %d", g.startedWith, players)
	}
}

func (g *SpiedGame) AssertFinish(t *testing.T, winner string) {
	t.Helper()
	if g.finishedWith != winner {
		t.Errorf("finish: got %q want %q", g.finishedWith, winner)
	}
}

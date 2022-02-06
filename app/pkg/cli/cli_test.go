package cli_test

import (
	"bytes"
	"jitsusama/lgwt/app/pkg/cli"
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	t.Run("records chris winning", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stdin := strings.NewReader("5\nChris wins\n")
		game := &spiedGame{}

		cli := cli.NewCli(stdin, stdout, game)
		cli.PlayPoker()

		if game.finishedWith != "Chris" {
			t.Errorf("end: got %q want %q", game.finishedWith, "Chris")
		}
	})
	t.Run("records cleo winning", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stdin := strings.NewReader("5\nCleo wins\n")
		game := &spiedGame{}

		cli := cli.NewCli(stdin, stdout, game)
		cli.PlayPoker()

		if game.finishedWith != "Cleo" {
			t.Errorf("end: got %q want %q", game.finishedWith, "Cleo")
		}
	})
	t.Run("requests number of players from user", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stdin := strings.NewReader("7\n")
		game := &spiedGame{}

		cli := cli.NewCli(stdin, stdout, game)
		cli.PlayPoker()

		actual := stdout.String()
		expected := "Please enter the number of players: "
		if actual != expected {
			t.Errorf("prompt: got %q want %q", actual, expected)
		}
		if game.startedWith != 7 {
			t.Errorf("start: want %d got %d", 7, game.startedWith)
		}
	})
}

type spiedGame struct {
	startedWith  int
	finishedWith string
}

func (g *spiedGame) Start(players int) {
	g.startedWith = players
}

func (g *spiedGame) Finish(winner string) {
	g.finishedWith = winner
}

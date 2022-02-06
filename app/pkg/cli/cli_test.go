package cli_test

import (
	"bytes"
	"jitsusama/lgwt/app/pkg/cli"
	"strings"
	"testing"
)

var (
	initialPrompt      = "Please enter the number of players: "
	invalidInputPrompt = "Bad value received for number of players, please try again with a number"
)

func TestCli(t *testing.T) {
	t.Run("records chris winning", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stdin := strings.NewReader("5\nChris wins\n")
		game := &spiedGame{}

		c := cli.NewCli(stdin, stdout, game)
		c.PlayGame()

		if game.finishedWith != "Chris" {
			t.Errorf("end: got %q want %q", game.finishedWith, "Chris")
		}
	})
	t.Run("records cleo winning", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stdin := strings.NewReader("5\nCleo wins\n")
		game := &spiedGame{}

		c := cli.NewCli(stdin, stdout, game)
		c.PlayGame()

		if game.finishedWith != "Cleo" {
			t.Errorf("end: got %q want %q", game.finishedWith, "Cleo")
		}
	})
	t.Run("requests number of players from user", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stdin := strings.NewReader("7\n")
		game := &spiedGame{}

		c := cli.NewCli(stdin, stdout, game)
		c.PlayGame()

		actual := stdout.String()
		if actual != initialPrompt {
			t.Errorf("prompt: got %q want %q", actual, initialPrompt)
		}
		if game.startedWith != 7 {
			t.Errorf("start: want %d got %d", 7, game.startedWith)
		}
	})
	t.Run("complains when non-numeric number is entered", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stdin := strings.NewReader("Pies\n")
		game := &spiedGame{}

		c := cli.NewCli(stdin, stdout, game)
		c.PlayGame()

		if game.started {
			t.Errorf("game was started")
		}

		actual := stdout.String()
		expected := initialPrompt + invalidInputPrompt
		if actual != expected {
			t.Errorf("got %q want %q", actual, expected)
		}
	})
}

type spiedGame struct {
	started      bool
	startedWith  int
	finishedWith string
}

func (g *spiedGame) Start(players int) {
	g.started = true
	g.startedWith = players
}

func (g *spiedGame) Finish(winner string) {
	g.finishedWith = winner
}

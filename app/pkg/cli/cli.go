package cli

import (
	"bufio"
	"io"
	"jitsusama/lgwt/app/pkg/storage"
	"strings"
)

func NewCli(store storage.PlayerStore, stdin io.Reader) *Cli {
	return &Cli{store, stdin}
}

type Cli struct {
	store storage.PlayerStore
	stdin io.Reader
}

func (c *Cli) PlayPoker() {
	line := c.readLine()
	player := parseLine(line)
	c.store.IncrementScore(player)
}

func (c *Cli) readLine() string {
	scanner := bufio.NewScanner(c.stdin)
	scanner.Scan()
	return scanner.Text()
}

func parseLine(line string) string {
	return strings.Replace(line, " wins", "", 1)
}

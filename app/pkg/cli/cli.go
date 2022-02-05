package cli

import (
	"bufio"
	"io"
	"jitsusama/lgwt/app/pkg/storage"
	"strings"
)

func NewCli(store storage.PlayerStore, stdin io.Reader) *Cli {
	return &Cli{store, bufio.NewScanner(stdin)}
}

type Cli struct {
	store storage.PlayerStore
	stdin *bufio.Scanner
}

func (c *Cli) PlayPoker() {
	line := c.readLine()
	player := parseLine(line)
	c.store.IncrementScore(player)
}

func (c *Cli) readLine() string {
	c.stdin.Scan()
	return c.stdin.Text()
}

func parseLine(line string) string {
	return strings.Replace(line, " wins", "", 1)
}

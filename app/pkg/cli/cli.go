package cli

import (
	"bufio"
	"io"
	"jitsusama/lgwt/app/pkg/storage"
	"strings"
	"time"
)

func NewCli(store storage.PlayerStore, stdin io.Reader, alerter BlindAlerter) *Cli {
	return &Cli{store, bufio.NewScanner(stdin), alerter}
}

type Cli struct {
	store storage.PlayerStore
	stdin *bufio.Scanner
	alert BlindAlerter
}

func (c *Cli) PlayPoker() {
	c.alert.ScheduleAlertAt(5*time.Second, 100)
	line := c.readLine()
	player := parseLine(line)
	c.store.IncrementScore(player)
}

func (c *Cli) readLine() string {
	c.stdin.Scan()
	return c.stdin.Text()
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

func parseLine(line string) string {
	return strings.Replace(line, " wins", "", 1)
}

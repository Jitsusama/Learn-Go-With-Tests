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
	c.scheduleBlindAlerts()
	c.waitForWin()
}

func (c *Cli) waitForWin() {
	line := c.readLine()
	player := parseLine(line)
	c.store.IncrementScore(player)
}

func (c *Cli) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		c.alert.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func (c *Cli) readLine() string {
	c.stdin.Scan()
	return c.stdin.Text()
}

func parseLine(line string) string {
	return strings.Replace(line, " wins", "", 1)
}

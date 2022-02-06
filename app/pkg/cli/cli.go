package cli

import (
	"bufio"
	"fmt"
	"io"
	"jitsusama/lgwt/app/pkg/game"
	"strconv"
	"strings"
)

func NewCli(stdin io.Reader, stdout io.Writer, game game.Game) *Cli {
	return &Cli{bufio.NewScanner(stdin), stdout, game}
}

type Cli struct {
	stdin  *bufio.Scanner
	stdout io.Writer
	game   game.Game
}

func (c *Cli) PlayGame() {
	players := c.getPlayerCount()

	c.game.Start(players)
	winner := c.waitForWin()
	c.game.Finish(winner)
}

func (c *Cli) getPlayerCount() int {
	fmt.Fprint(c.stdout, "Please enter the number of players: ")
	players, _ := strconv.Atoi(c.readLine())
	return players
}

func (c *Cli) waitForWin() string {
	line := c.readLine()
	return parseLine(line)
}

func (c *Cli) readLine() string {
	c.stdin.Scan()
	return c.stdin.Text()
}

func parseLine(line string) string {
	return strings.Replace(line, " wins", "", 1)
}

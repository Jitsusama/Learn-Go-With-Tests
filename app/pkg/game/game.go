package game

import "io"

type Game interface {
	Start(players int, alertDestination io.Writer)
	Finish(winner string)
}

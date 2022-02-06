package game

type Game interface {
	Start(players int)
	Finish(winner string)
}

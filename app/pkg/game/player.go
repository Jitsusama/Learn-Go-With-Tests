package game

func NewPlayer(name string, wins int) *Player {
	return &Player{name, wins}
}

type Player struct {
	Name string
	Wins int
}

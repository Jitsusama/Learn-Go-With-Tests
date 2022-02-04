package storage

type PlayerStore interface {
	GetScore(name string) int
	IncrementScore(name string)
	GetLeague() []Player
}

type Player struct {
	Name string
	Wins int
}

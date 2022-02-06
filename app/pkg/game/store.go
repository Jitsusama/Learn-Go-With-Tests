package game

type PlayerStore interface {
	GetScore(name string) int
	IncrementScore(name string)
	GetLeague() League
}

package storage

type PlayerStore interface {
	GetScore(name string) int
	IncrementScore(name string)
	GetLeague() League
}

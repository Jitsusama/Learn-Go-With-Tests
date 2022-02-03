package server

func NewPlayerStoreInMemory() *PlayerStoreInMemory {
	return &PlayerStoreInMemory{map[string]int{}}
}

type PlayerStoreInMemory struct {
	store map[string]int
}

func (s *PlayerStoreInMemory) GetScore(name string) int {
	return s.store[name]
}

func (s *PlayerStoreInMemory) IncrementScore(name string) {
	s.store[name]++
}

func (s *PlayerStoreInMemory) GetLeague() []Player {
	return nil
}

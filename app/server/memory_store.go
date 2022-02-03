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
	var league []Player
	for name, wins := range s.store {
		league = append(league, Player{name, wins})
	}
	return league
}

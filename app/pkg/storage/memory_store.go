package storage

import "jitsusama/lgwt/app/pkg/game"

func NewMemoryPlayerStore() *MemoryPlayerStore {
	return &MemoryPlayerStore{map[string]int{}}
}

type MemoryPlayerStore struct {
	store map[string]int
}

func (s *MemoryPlayerStore) GetScore(name string) int {
	return s.store[name]
}

func (s *MemoryPlayerStore) IncrementScore(name string) {
	s.store[name]++
}

func (s *MemoryPlayerStore) GetLeague() game.League {
	var league game.League
	for name, wins := range s.store {
		league = append(league, game.Player{
			Name: name, Wins: wins,
		})
	}
	return league
}

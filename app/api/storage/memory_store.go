package storage

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

func (s *MemoryPlayerStore) GetLeague() League {
	var league League
	for name, wins := range s.store {
		league = append(league, Player{
			Name: name, Wins: wins,
		})
	}
	return league
}

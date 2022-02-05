package testing

import (
	"jitsusama/lgwt/app/pkg/storage"
	"testing"
)

func NewStubbedPlayerStore(scores map[string]int, league storage.League) *StubbedPlayerStore {
	return &StubbedPlayerStore{scores, nil, league}
}

type StubbedPlayerStore struct {
	scores map[string]int
	wins   []string
	league storage.League
}

func (s *StubbedPlayerStore) GetScore(name string) int {
	return s.scores[name]
}

func (s *StubbedPlayerStore) IncrementScore(name string) {
	s.wins = append(s.wins, name)
}

func (s *StubbedPlayerStore) GetLeague() storage.League {
	return s.league
}

func AssertPlayerWin(t testing.TB, store *StubbedPlayerStore, player string) {
	if len(store.wins) != 1 {
		t.Fatalf("calls: got %d want %d", len(store.wins), 1)
	}
	if store.wins[0] != player {
		t.Errorf("player: got %q want %q", store.wins[0], player)
	}
}

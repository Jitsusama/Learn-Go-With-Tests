package cli_test

import (
	"jitsusama/lgwt/app/pkg/cli"
	"jitsusama/lgwt/app/pkg/storage"
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	t.Run("records chris winning", func(t *testing.T) {
		stdin := strings.NewReader("Chris wins\n")
		store := StubPlayerStore{}

		cli := cli.NewCli(&store, stdin)
		cli.PlayPoker()

		assertPlayerWin(t, &store, "Chris")
	})
	t.Run("records cleo winning", func(t *testing.T) {
		stdin := strings.NewReader("Cleo wins\n")
		store := StubPlayerStore{}

		cli := cli.NewCli(&store, stdin)
		cli.PlayPoker()

		assertPlayerWin(t, &store, "Cleo")
	})
}

type StubPlayerStore struct {
	scores map[string]int
	wins   []string
	league storage.League
}

func (s *StubPlayerStore) GetScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) IncrementScore(name string) {
	s.wins = append(s.wins, name)
}

func (s *StubPlayerStore) GetLeague() storage.League {
	return s.league
}

func assertPlayerWin(t testing.TB, store *StubPlayerStore, player string) {
	if len(store.wins) != 1 {
		t.Fatalf("calls: got %d want %d", len(store.wins), 1)
	}
	if store.wins[0] != player {
		t.Errorf("player: got %q want %q", store.wins[0], player)
	}
}

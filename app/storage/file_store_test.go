package storage

import (
	"reflect"
	"strings"
	"testing"
)

func TestFileStorage(t *testing.T) {
	t.Run("dummy test", func(t *testing.T) {
		fileContents := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		fileStore := FilePlayerStore{fileContents}

		// read once
		actual := fileStore.GetLeague()
		expected := []Player{{"Cleo", 10}, {"Chris", 33}}
		assertLeagueBody(t, actual, expected)

		// read once more
		actual = fileStore.GetLeague()
		expected = []Player{{"Cleo", 10}, {"Chris", 33}}
		assertLeagueBody(t, actual, expected)
	})
}

func assertLeagueBody(t testing.TB, actual []Player, expected []Player) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

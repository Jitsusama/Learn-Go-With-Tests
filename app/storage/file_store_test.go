package storage

import (
	"reflect"
	"strings"
	"testing"
)

func TestFileStorage(t *testing.T) {
	t.Run("reads league data from JSON file", func(t *testing.T) {
		fileContents := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		fileStore := FilePlayerStore{fileContents}

		// read once
		actual := fileStore.GetLeague()
		expected := []Player{{"Cleo", 10}, {"Chris", 33}}
		assertLeague(t, actual, expected)

		// read once more
		actual = fileStore.GetLeague()
		expected = []Player{{"Cleo", 10}, {"Chris", 33}}
		assertLeague(t, actual, expected)
	})

	t.Run("reads player score from JSON file", func(t *testing.T) {
		fileContents := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		fileStore := FilePlayerStore{fileContents}

		actual := fileStore.GetPlayerScore("Chris")
		expected := 33
		assertScore(t, actual, expected)
	})
}

func assertLeague(t testing.TB, actual []Player, expected []Player) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func assertScore(t testing.TB, actual int, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

package storage_test

import (
	"io/ioutil"
	"jitsusama/lgwt/app/pkg/storage"
	"os"
	"reflect"
	"testing"
)

func TestFileStorage(t *testing.T) {
	t.Run("reads sorted league data from JSON file", func(t *testing.T) {
		file, cleanup := createFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanup()
		store, err := storage.NewFilePlayerStore(file)

		// read once
		actual := store.GetLeague()
		expected := storage.League{{"Chris", 33}, {"Cleo", 10}}
		assertLeague(t, actual, expected)

		// read once more
		actual = store.GetLeague()
		expected = storage.League{{"Chris", 33}, {"Cleo", 10}}
		assertLeague(t, actual, expected)
		assertNoError(t, err)
	})

	t.Run("reads player score from JSON file", func(t *testing.T) {
		file, cleanup := createFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
			]`)
		defer cleanup()
		store, err := storage.NewFilePlayerStore(file)

		actual := store.GetScore("Chris")
		expected := 33
		assertScore(t, actual, expected)
		assertNoError(t, err)
	})

	t.Run("increments score for existing player", func(t *testing.T) {
		file, cleanup := createFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanup()
		store, err := storage.NewFilePlayerStore(file)

		store.IncrementScore("Chris")

		actual := store.GetScore("Chris")
		expected := 34
		assertScore(t, actual, expected)
		assertNoError(t, err)
	})

	t.Run("stores score for new player", func(t *testing.T) {
		file, cleanup := createFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanup()
		store, err := storage.NewFilePlayerStore(file)

		store.IncrementScore("Pepper")

		actual := store.GetScore("Pepper")
		expected := 1
		assertScore(t, actual, expected)
		assertNoError(t, err)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		file, cleanup := createFile(t, "")
		defer cleanup()

		_, err := storage.NewFilePlayerStore(file)

		assertNoError(t, err)
	})
}

func assertLeague(t testing.TB, actual storage.League, expected storage.League) {
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

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got an error: %v", err)
	}
}

func createFile(t testing.TB, contents string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "*.json")
	if err != nil {
		t.Fatalf("temp file creation error: %v", err)
	}
	tmpFile.Write([]byte(contents))

	return tmpFile, func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}
}

package storage_test

import (
	"io"
	"io/ioutil"
	"jitsusama/lgwt/app/storage"
	"os"
	"reflect"
	"testing"
)

func TestFileStorage(t *testing.T) {
	t.Run("reads league data from JSON file", func(t *testing.T) {
		file, cleanup := createFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanup()
		store := storage.NewFilePlayerStore(file)

		// read once
		actual := store.GetLeague()
		expected := storage.League{{"Cleo", 10}, {"Chris", 33}}
		assertLeague(t, actual, expected)

		// read once more
		actual = store.GetLeague()
		expected = storage.League{{"Cleo", 10}, {"Chris", 33}}
		assertLeague(t, actual, expected)
	})

	t.Run("reads player score from JSON file", func(t *testing.T) {
		file, cleanup := createFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanup()
		store := storage.NewFilePlayerStore(file)

		actual := store.GetScore("Chris")
		expected := 33
		assertScore(t, actual, expected)
	})

	t.Run("increments score for existing player", func(t *testing.T) {
		file, cleanup := createFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanup()
		store := storage.NewFilePlayerStore(file)

		store.IncrementScore("Chris")

		actual := store.GetScore("Chris")
		expected := 34
		assertScore(t, actual, expected)
	})

	t.Run("stores score for new player", func(t *testing.T) {
		file, cleanup := createFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanup()
		store := storage.NewFilePlayerStore(file)

		store.IncrementScore("Pepper")

		actual := store.GetScore("Pepper")
		expected := 1
		assertScore(t, actual, expected)
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

func createFile(t testing.TB, contents string) (io.ReadWriteSeeker, func()) {
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

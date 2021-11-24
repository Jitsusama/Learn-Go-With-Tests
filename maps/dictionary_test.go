package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("finds a known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")

		assertStrings(t, got, "this is just a test")
	})

	t.Run("complains on unknown words", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

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

func TestAdd(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	t.Run("add a new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("complains when redefining a word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	originalDefinition := "this is just a test"
	updatedDefinition := "new definition"

	t.Run("updates an existing word", func(t *testing.T) {
		dictionary := Dictionary{word: originalDefinition}

		err := dictionary.Update(word, updatedDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, updatedDefinition)
	})

	t.Run("complains when word does not exist", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update(word, originalDefinition)

		assertError(t, err, ErrWordDoesNotExist)
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

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("unable to find word:", err)
	} else if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

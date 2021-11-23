package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Joel", "")
		want := "Hello, Joel"
		assertMessage(t, got, want)
	})

	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)
	})

	t.Run("says hello to a person in spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan"
		assertMessage(t, got, want)
	})
}

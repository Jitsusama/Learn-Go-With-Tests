package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns fastest server", func(t *testing.T) {
		slowServer := makeDelayedServer(time.Millisecond * 20)
		defer slowServer.Close()

		fastServer := makeDelayedServer(time.Millisecond * 0)
		defer fastServer.Close()

		got, _ := Racer(slowServer.URL, fastServer.URL)

		if got != fastServer.URL {
			t.Errorf("got %q, want %q", got, fastServer.URL)
		}
	})
	t.Run("complains if neither server responds within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		defer serverA.Close()
		serverB := makeDelayedServer(12 * time.Second)
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Errorf("expected an error, but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

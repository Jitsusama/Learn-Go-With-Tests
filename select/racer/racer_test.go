package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(time.Millisecond * 20)
	defer slowServer.Close()

	fastServer := makeDelayedServer(time.Millisecond * 0)
	defer fastServer.Close()

	got := Racer(slowServer.URL, fastServer.URL)

	if got != fastServer.URL {
		t.Errorf("got %q, want %q", got, fastServer.URL)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

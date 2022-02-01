package server_test

import (
	"bytes"
	"fmt"
	"jitsusama/lgwt/app/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerRetrieval(t *testing.T) {
	store := StubPlayerStore{map[string]int{"Pepper": 20, "Floyd": 10}}
	server := &server.PlayerServer{&store}

	t.Run("retrieve pepper's score", func(t *testing.T) {
		request := getScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertBody(t, response.Body, "20")
	})
	t.Run("retrieve floyd's score", func(t *testing.T) {
		request := getScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertBody(t, response.Body, "10")
	})
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetScore(name string) int {
	return s.scores[name]
}

func getScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/players/%s", player), nil)
	return req
}

func assertBody(t testing.TB, body *bytes.Buffer, expected string) {
	t.Helper()
	actual := body.String()
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}

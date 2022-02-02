package server_test

import (
	"bytes"
	"fmt"
	"jitsusama/lgwt/app/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScoreRetrieval(t *testing.T) {
	store := StubPlayerStore{map[string]int{"Pepper": 20, "Floyd": 10}, nil}
	server := server.NewPlayerServer(&store)

	t.Run("retrieve pepper's score", func(t *testing.T) {
		request := getPlayer("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, 200)
		assertBody(t, response.Body, "20")
	})
	t.Run("retrieve floyd's score", func(t *testing.T) {
		request := getPlayer("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, 200)
		assertBody(t, response.Body, "10")
	})
	t.Run("complains on missing players", func(t *testing.T) {
		request := getPlayer("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, 404)
	})
}

func TestScoreStorage(t *testing.T) {
	store := StubPlayerStore{map[string]int{}, nil}
	server := server.NewPlayerServer(&store)

	t.Run("records scores", func(t *testing.T) {
		player := "Pepper"
		request := postPlayer(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, 202)
		if len(store.posts) != 1 {
			t.Errorf("got %d want %d", len(store.posts), 1)
		}
		if store.posts[0] != player {
			t.Errorf("got %q want %q", store.posts[0], player)
		}
	})
}

func TestLeagueRetrieval(t *testing.T) {
	store := StubPlayerStore{}
	server := server.NewPlayerServer(&store)

	t.Run("stupid test", func(t *testing.T) {
		request := getLeague()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, 200)
	})
}

type StubPlayerStore struct {
	scores map[string]int
	posts  []string
}

func (s *StubPlayerStore) GetScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) IncrementScore(name string) {
	s.posts = append(s.posts, name)
}

func getPlayer(player string) *http.Request {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/players/%s", player), nil)
	return req
}

func postPlayer(player string) *http.Request {
	req, _ := http.NewRequest("POST", fmt.Sprintf("/players/%s", player), nil)
	return req
}

func getLeague() *http.Request {
	req, _ := http.NewRequest("GET", "/league", nil)
	return req
}

func assertBody(t testing.TB, body *bytes.Buffer, expected string) {
	t.Helper()
	actual := body.String()
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}

func assertStatus(t testing.TB, actual int, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %d want %d", actual, expected)
	}
}

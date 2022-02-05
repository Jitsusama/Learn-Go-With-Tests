package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jitsusama/lgwt/app/pkg/server"
	"jitsusama/lgwt/app/pkg/storage"
	test "jitsusama/lgwt/app/pkg/testing"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestScoreRetrieval(t *testing.T) {
	store := test.NewStubbedPlayerStore(map[string]int{"Pepper": 20, "Floyd": 10}, nil)
	server := server.NewPlayerServer(store)

	t.Run("retrieve pepper's score", func(t *testing.T) {
		request := getPlayer("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, 200)
		assertPlayerBody(t, response.Body, "20")
	})
	t.Run("retrieve floyd's score", func(t *testing.T) {
		request := getPlayer("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, 200)
		assertPlayerBody(t, response.Body, "10")
	})
	t.Run("complains on missing players", func(t *testing.T) {
		request := getPlayer("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, 404)
	})
}

func TestScoreStorage(t *testing.T) {
	store := &test.StubbedPlayerStore{}
	server := server.NewPlayerServer(store)

	t.Run("records scores", func(t *testing.T) {
		player := "Pepper"
		request := postPlayer(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, 202)
		test.AssertPlayerWin(t, store, player)
	})
}

func TestLeagueRetrieval(t *testing.T) {
	t.Run("retrieves scores of entire league", func(t *testing.T) {
		expected := []storage.Player{
			{Name: "Cleo", Wins: 32}, {Name: "Chris", Wins: 20},
			{Name: "Test", Wins: 14},
		}
		store := test.NewStubbedPlayerStore(nil, expected)
		server := server.NewPlayerServer(store)

		request := getLeague()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, 200)
		assertContentType(t, response.Result().Header, "application/json")
		assertLeagueBody(t, response.Body, expected)
	})
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

func assertStatus(t testing.TB, actual int, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %d want %d", actual, expected)
	}
}

func assertContentType(t testing.TB, headers http.Header, expected string) {
	t.Helper()
	actual := headers.Get("content-type")
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func assertPlayerBody(t testing.TB, body *bytes.Buffer, expected string) {
	t.Helper()
	actual := body.String()
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}

func assertLeagueBody(t testing.TB, body *bytes.Buffer, expected []storage.Player) {
	t.Helper()
	var actual []storage.Player
	if err := json.NewDecoder(body).Decode(&actual); err != nil {
		t.Fatalf("unable to parse %q: '%v'", body, err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

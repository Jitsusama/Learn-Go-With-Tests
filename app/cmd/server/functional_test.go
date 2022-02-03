package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jitsusama/lgwt/app/server"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRecordsTotalWinsAndAllowsTotalRetrieval(t *testing.T) {
	player := "Pepper"

	str := server.NewPlayerStoreInMemory()
	svr := server.NewPlayerServer(str)

	svr.ServeHTTP(httptest.NewRecorder(), postPlayer(player))
	svr.ServeHTTP(httptest.NewRecorder(), postPlayer(player))
	svr.ServeHTTP(httptest.NewRecorder(), postPlayer(player))

	t.Run("associates wins with player", func(t *testing.T) {
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, getPlayer(player))

		assertStatus(t, response.Code, 200)
		assertPlayerBody(t, response.Body, "3")
	})

	t.Run("associates player with league", func(t *testing.T) {
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, getLeague())

		assertStatus(t, response.Code, 200)
		assertLeagueBody(t, response.Body, []server.Player{
			{Name: "Pepper", Wins: 3},
		})
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

func assertPlayerBody(t testing.TB, body *bytes.Buffer, expected string) {
	t.Helper()
	actual := body.String()
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}

func assertLeagueBody(t testing.TB, body *bytes.Buffer, expected []server.Player) {
	t.Helper()
	var actual []server.Player
	if err := json.NewDecoder(body).Decode(&actual); err != nil {
		t.Fatalf("unable to parse %q: '%v'", body, err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

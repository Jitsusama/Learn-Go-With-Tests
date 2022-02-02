package main

import (
	"bytes"
	"fmt"
	"jitsusama/lgwt/app/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordsTotalWinsAndAllowsTotalRetrieval(t *testing.T) {
	player := "Pepper"
	response := httptest.NewRecorder()

	store := server.NewPlayerStoreInMemory()
	server := server.NewPlayerServer(store)

	server.ServeHTTP(httptest.NewRecorder(), postPlayer(player))
	server.ServeHTTP(httptest.NewRecorder(), postPlayer(player))
	server.ServeHTTP(httptest.NewRecorder(), postPlayer(player))
	server.ServeHTTP(response, getPlayer(player))

	assertStatus(t, response.Code, 200)
	assertBody(t, response.Body, "3")
}

func getPlayer(player string) *http.Request {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/players/%s", player), nil)
	return req
}

func postPlayer(player string) *http.Request {
	req, _ := http.NewRequest("POST", fmt.Sprintf("/players/%s", player), nil)
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

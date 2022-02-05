package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jitsusama/lgwt/app/api/server"
	"jitsusama/lgwt/app/api/storage"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestRecordsTotalWinsAndAllowsTotalRetrieval(t *testing.T) {
	player := "Pepper"

	file, cleanup := createFile(t, "[]")
	defer cleanup()
	str, _ := storage.NewFilePlayerStore(file)
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
		assertLeagueBody(t, response.Body, []storage.Player{
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

func createFile(t testing.TB, contents string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "*.json")
	if err != nil {
		t.Fatalf("temp file creation error: %v", err)
	}
	tmpFile.Write([]byte(contents))

	return tmpFile, func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}
}

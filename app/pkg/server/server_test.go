package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jitsusama/lgwt/app/pkg/game"
	"jitsusama/lgwt/app/pkg/server"
	test "jitsusama/lgwt/app/pkg/testing"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestScoreRetrieval(t *testing.T) {
	store := test.NewStubbedPlayerStore(map[string]int{"Pepper": 20, "Floyd": 10}, nil)
	server := makeServer(t, store, &test.SpiedGame{})

	t.Run("retrieve pepper's score", func(t *testing.T) {
		request := getPlayer("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
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

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestScoreStorage(t *testing.T) {
	store := &test.StubbedPlayerStore{}
	server := makeServer(t, store, &test.SpiedGame{})

	t.Run("records scores", func(t *testing.T) {
		player := "Pepper"
		request := postPlayer(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		test.AssertPlayerWin(t, store, player)
	})
}

func TestLeagueRetrieval(t *testing.T) {
	t.Run("retrieves scores of entire league", func(t *testing.T) {
		expected := []game.Player{
			{Name: "Cleo", Wins: 32}, {Name: "Chris", Wins: 20},
			{Name: "Test", Wins: 14},
		}
		store := test.NewStubbedPlayerStore(nil, expected)
		server := makeServer(t, store, &test.SpiedGame{})

		request := getLeague()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response.Result().Header, "application/json")
		assertLeagueBody(t, response.Body, expected)
	})
}

func TestGame(t *testing.T) {
	t.Run("presents game HTML page", func(t *testing.T) {
		server := makeServer(t, &test.StubbedPlayerStore{}, &test.SpiedGame{})

		req, _ := http.NewRequest("GET", "/game", nil)
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusOK)
	})
	t.Run("handles a 3 player game, sending blind update messages and allowing Ruth to be declared the winner", func(t *testing.T) {
		store := &test.StubbedPlayerStore{}
		game := test.NewSpiedGame([]byte("Blind is 100"))
		server := httptest.NewServer(makeServer(t, store, game))
		defer server.Close()

		ws := getWebSocket(t, server.URL)
		defer ws.Close()

		sendToWebSocket(t, ws, "3")
		sendToWebSocket(t, ws, "Ruth")

		game.AssertStart(t, 3)
		game.AssertFinish(t, "Ruth")
		within(t, 10*time.Millisecond, func() {
			_, actual, _ := ws.ReadMessage()
			if string(actual) != "Blind is 100" {
				t.Errorf("blind: got %q want %q", string(actual), "Blind is 100")
			}
		})
	})
}

func makeServer(t *testing.T, store game.PlayerStore, game game.Game) *server.PlayerServer {
	t.Helper()
	server, err := server.NewPlayerServer(store, game)
	if err != nil {
		t.Fatalf("server creation failure: %v", err)
	}
	return server
}

func getWebSocket(t *testing.T, url string) *websocket.Conn {
	uri := fmt.Sprintf("ws://%s/ws", strings.TrimPrefix(url, "http://"))
	ws, _, err := websocket.DefaultDialer.Dial(uri, nil)
	if err != nil {
		t.Fatalf("socket creation failure: %v", err)
	}
	return ws
}

func sendToWebSocket(t *testing.T, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("message sending failure: %v", err)
	}
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

func assertLeagueBody(t testing.TB, body *bytes.Buffer, expected []game.Player) {
	t.Helper()
	var actual []game.Player
	if err := json.NewDecoder(body).Decode(&actual); err != nil {
		t.Fatalf("unable to parse %q: '%v'", body, err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)
	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

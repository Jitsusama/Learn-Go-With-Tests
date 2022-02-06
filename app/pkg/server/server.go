package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"jitsusama/lgwt/app/pkg/game"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var (
	//go:embed index.html
	templates embed.FS
)

type PlayerServer struct {
	store  game.PlayerStore
	router *http.ServeMux
	game   game.Game
	tmpl   *template.Template
}

func NewPlayerServer(store game.PlayerStore, game game.Game) (*PlayerServer, error) {
	tmpl, err := template.ParseFS(templates, "index.html")
	if err != nil {
		return nil, fmt.Errorf("problem opening templates: %v", err)
	}

	p := &PlayerServer{store, http.NewServeMux(), game, tmpl}
	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	p.router.Handle("/game", http.HandlerFunc(p.gameHandler))
	p.router.Handle("/ws", http.HandlerFunc(p.webSocketHandler))
	return p, nil
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case "POST":
		p.scoreIncrease(w, player)
	case "GET":
		p.scoreRetrieval(w, player)
	}
}

func (p *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	p.tmpl.Execute(w, nil)
}

func (p *PlayerServer) webSocketHandler(w http.ResponseWriter, r *http.Request) {
	s := newIncomingSocket(w, r)

	playersMessage := s.WaitForMessage()
	players, _ := strconv.Atoi(playersMessage)
	p.game.Start(players, s)

	winnerMessage := s.WaitForMessage()
	p.game.Finish(winnerMessage)
}

func (p *PlayerServer) scoreRetrieval(w http.ResponseWriter, player string) {
	score := p.store.GetScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) scoreIncrease(w http.ResponseWriter, player string) {
	p.store.IncrementScore(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

package server

import (
	"encoding/json"
	"fmt"
	"jitsusama/lgwt/app/pkg/storage"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store  storage.PlayerStore
	router *http.ServeMux
}

func NewPlayerServer(store storage.PlayerStore) *PlayerServer {
	p := &PlayerServer{store, http.NewServeMux()}
	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(200)
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

func (p *PlayerServer) scoreRetrieval(w http.ResponseWriter, player string) {
	score := p.store.GetScore(player)
	if score == 0 {
		w.WriteHeader(404)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) scoreIncrease(w http.ResponseWriter, player string) {
	p.store.IncrementScore(player)
	w.WriteHeader(202)
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

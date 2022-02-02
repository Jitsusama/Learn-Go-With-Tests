package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetScore(name string) int
	IncrementScore(name string)
}

type PlayerServer struct {
	store  PlayerStore
	router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{store, http.NewServeMux()}
	p.router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	p.router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		player := strings.TrimPrefix(r.URL.Path, "/players/")

		switch r.Method {
		case "POST":
			p.scoreIncrease(w, player)
		case "GET":
			p.scoreRetrieval(w, player)
		}
	}))
	return p
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
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

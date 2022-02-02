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
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case "POST":
		p.scoreIncrease(w, player)
	case "GET":
		p.scoreRetrieval(w, player)
	}
}

func (p *PlayerServer) scoreRetrieval(w http.ResponseWriter, player string) {
	score := p.Store.GetScore(player)

	if score == 0 {
		w.WriteHeader(404)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) scoreIncrease(w http.ResponseWriter, player string) {
	p.Store.IncrementScore(player)
	w.WriteHeader(202)
}

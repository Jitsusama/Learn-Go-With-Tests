package main

import (
	"jitsusama/lgwt/app/server"
	"log"
	"net/http"
)

type PlayerStoreInMemory struct{}

func (s *PlayerStoreInMemory) GetScore(name string) int {
	return 123
}

func main() {
	server := &server.PlayerServer{Store: &PlayerStoreInMemory{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}

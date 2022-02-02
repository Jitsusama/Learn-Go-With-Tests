package main

import (
	"jitsusama/lgwt/app/server"
	"log"
	"net/http"
)

func main() {
	store := server.PlayerStoreInMemory{}
	server := &server.PlayerServer{Store: &store}
	log.Fatal(http.ListenAndServe(":5000", server))
}

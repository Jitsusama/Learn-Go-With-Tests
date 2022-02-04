package main

import (
	"jitsusama/lgwt/app/server"
	"jitsusama/lgwt/app/storage"
	"log"
	"net/http"
)

func main() {
	store := storage.PlayerStoreInMemory{}
	server := server.NewPlayerServer(&store)
	log.Fatal(http.ListenAndServe(":5000", server))
}

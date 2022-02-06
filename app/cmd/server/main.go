package main

import (
	"jitsusama/lgwt/app/pkg/server"
	"jitsusama/lgwt/app/pkg/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.OpenFile("game.db.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %q: %v", "game.db.json", err)
	}
	store, err := storage.NewFilePlayerStore(file)
	if err != nil {
		log.Fatalf("problem creating store: %v", err)
	}

	server, err := server.NewPlayerServer(store)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

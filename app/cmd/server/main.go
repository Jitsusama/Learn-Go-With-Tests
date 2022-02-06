package main

import (
	"jitsusama/lgwt/app/pkg/game"
	"jitsusama/lgwt/app/pkg/server"
	"jitsusama/lgwt/app/pkg/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	store := createGameStore("game.db.json")
	game := createGame(store)
	server := createServer(store, game)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func createGameStore(filename string) game.PlayerStore {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %q: %v", filename, err)
	}

	store, err := storage.NewFilePlayerStore(file)
	if err != nil {
		log.Fatalf("problem creating store: %v", err)
	}
	return store
}

func createGame(store game.PlayerStore) game.Game {
	alerter := game.BlindAlerterFunc(game.GenericAlerter)
	return game.NewPokerGame(alerter, store)
}

func createServer(store game.PlayerStore, game game.Game) *server.PlayerServer {
	server, err := server.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	return server
}

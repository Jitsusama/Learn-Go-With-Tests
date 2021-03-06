package main

import (
	"fmt"
	"jitsusama/lgwt/app/pkg/cli"
	"jitsusama/lgwt/app/pkg/game"
	"jitsusama/lgwt/app/pkg/storage"
	"log"
	"os"
)

func main() {
	fmt.Println("Type `{Name} wins<CR>` to record a win.")

	file, err := os.OpenFile("game.db.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %q: %v", file.Name(), err)
	}
	store, err := storage.NewFilePlayerStore(file)
	if err != nil {
		log.Fatalf("problem creating store: %v", err)
	}

	alerter := game.BlindAlerterFunc(game.GenericAlerter)
	game := game.NewPokerGame(alerter, store)
	c := cli.NewCli(os.Stdin, os.Stdout, game)
	c.PlayGame()
}

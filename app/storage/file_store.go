package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func NewFilePlayerStore(database io.ReadWriteSeeker) *FilePlayerStore {
	database.Seek(0, 0)

	var league League
	if err := json.NewDecoder(database).Decode(&league); err != nil {
		fmt.Printf("json parsing error: %v", err)
	}
	if database, ok := database.(*os.File); ok {
		return &FilePlayerStore{
			json.NewEncoder(&tape{database}),
			league,
		}
	}
	return nil
}

type FilePlayerStore struct {
	database *json.Encoder
	league   League
}

func (f *FilePlayerStore) GetScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FilePlayerStore) IncrementScore(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	f.database.Encode(f.league)
}

func (f *FilePlayerStore) GetLeague() League {
	return f.league
}

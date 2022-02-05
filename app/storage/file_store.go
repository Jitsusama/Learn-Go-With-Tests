package storage

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewFilePlayerStore(database io.ReadWriteSeeker) *FilePlayerStore {
	database.Seek(0, 0)

	var league League
	if err := json.NewDecoder(database).Decode(&league); err != nil {
		fmt.Printf("json parsing error: %v", err)
	}

	return &FilePlayerStore{&tape{database}, league}
}

type FilePlayerStore struct {
	database io.Writer
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
	// reset position to beginning of file
	json.NewEncoder(f.database).Encode(f.league)
}

func (f *FilePlayerStore) GetLeague() League {
	return f.league
}

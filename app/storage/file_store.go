package storage

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewFilePlayerStore(database io.ReadWriteSeeker) *FilePlayerStore {
	return &FilePlayerStore{database}
}

type FilePlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FilePlayerStore) GetScore(name string) int {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FilePlayerStore) IncrementScore(name string) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}
	// reset position to beginning of file
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

func (f *FilePlayerStore) GetLeague() League {
	// reset position to beginning of file
	f.database.Seek(0, 0)

	var league League
	if err := json.NewDecoder(f.database).Decode(&league); err != nil {
		fmt.Printf("json parsing error: %v", err)
	}
	return league
}

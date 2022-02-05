package storage

import (
	"encoding/json"
	"fmt"
	"io"
)

type FilePlayerStore struct {
	database io.ReadWriteSeeker
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

func (f *FilePlayerStore) GetPlayerScore(name string) int {
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
	}
	// reset position to beginning of file
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

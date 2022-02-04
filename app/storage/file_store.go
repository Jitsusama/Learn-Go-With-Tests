package storage

import (
	"encoding/json"
	"fmt"
	"io"
)

type FilePlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FilePlayerStore) GetLeague() []Player {
	// reset position to beginning of file
	f.database.Seek(0, 0)

	var league []Player
	if err := json.NewDecoder(f.database).Decode(&league); err != nil {
		fmt.Printf("json parsing error: %v", err)
	}
	return league
}

func (f *FilePlayerStore) GetPlayerScore(name string) int {
	for _, player := range f.GetLeague() {
		if player.Name == name {
			return player.Wins
		}
	}
	return 0
}

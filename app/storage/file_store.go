package storage

import (
	"encoding/json"
	"fmt"
	"io"
)

type FilePlayerStore struct {
	database io.ReadSeeker
}

func (f *FilePlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	var league []Player
	if err := json.NewDecoder(f.database).Decode(&league); err != nil {
		fmt.Printf("json parsing error: %v", err)
	}
	return league
}

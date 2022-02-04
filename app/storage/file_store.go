package storage

import (
	"encoding/json"
	"fmt"
	"io"
)

type FilePlayerStore struct {
	database io.Reader
}

func (f *FilePlayerStore) GetLeague() []Player {
	var league []Player
	if err := json.NewDecoder(f.database).Decode(&league); err != nil {
		fmt.Printf("json parsing error: %v", err)
	}
	return league
}

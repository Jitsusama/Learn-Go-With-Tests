package storage

import (
	"encoding/json"
	"fmt"
	"jitsusama/lgwt/app/pkg/game"
	"os"
	"sort"
)

func NewFilePlayerStore(database *os.File) (*FilePlayerStore, error) {
	err := sanitizeDatabase(database)
	if err != nil {
		return nil, err
	}
	league, err := decodeDatabase(database)
	if err != nil {
		return nil, err
	}
	return &FilePlayerStore{
		json.NewEncoder(&tape{database}),
		league,
	}, nil
}

type FilePlayerStore struct {
	database *json.Encoder
	league   game.League
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
		f.league = append(f.league, *game.NewPlayer(name, 1))
	}
	f.database.Encode(f.league)
}

func (f *FilePlayerStore) GetLeague() game.League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func sanitizeDatabase(database *os.File) error {
	database.Seek(0, 0)
	stat, err := database.Stat()
	if err != nil {
		return fmt.Errorf("problem reading file %q: %v", database.Name(), err)
	}
	// make empty file into valid empty league JSON
	if stat.Size() == 0 {
		database.Write([]byte("[]"))
		database.Seek(0, 0)
	}
	return nil
}

func decodeDatabase(database *os.File) (game.League, error) {
	var league game.League
	if err := json.NewDecoder(database).Decode(&league); err != nil {
		return nil, fmt.Errorf("json parsing error while reading %q: %v", database.Name(), err)
	}
	return league, nil
}

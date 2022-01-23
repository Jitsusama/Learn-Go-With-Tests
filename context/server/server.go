package server

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if data, err := store.Fetch(req.Context()); err == nil {
			fmt.Fprint(resp, data)
		} else {
			return
		}
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

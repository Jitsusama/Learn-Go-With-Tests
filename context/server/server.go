package server

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		data := make(chan string, 1)

		go func() { data <- store.Fetch() }()

		select {
		case data := <-data:
			fmt.Fprint(resp, data)
		case <-req.Context().Done():
			store.Cancel()
		}
	}
}

type Store interface {
	Fetch() string
	Cancel()
}

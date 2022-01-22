package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() { s.cancelled = true }

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		Server(store).ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q want %q", response.Body.String(), data)
		}
		if store.cancelled {
			t.Errorf("it should not be cancelled")
		}
	})

	t.Run("tells store to cancel work on request cancel", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data}

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		requestContext, cancel := context.WithCancel(request.Context())
		request = request.WithContext(requestContext)
		response := httptest.NewRecorder()

		time.AfterFunc(5*time.Millisecond, cancel)
		Server(store).ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
}

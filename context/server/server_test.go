package server

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() { s.cancelled = true }

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not cancelled")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was cancelled")
	}
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}

		request := httptest.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()

		Server(store).ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q want %q", response.Body.String(), data)
		}
		store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work on request cancel", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data, t: t}

		request := httptest.NewRequest("GET", "/", nil)
		requestContext, cancel := context.WithCancel(request.Context())
		request = request.WithContext(requestContext)
		response := httptest.NewRecorder()

		time.AfterFunc(5*time.Millisecond, cancel)
		Server(store).ServeHTTP(response, request)

		store.assertWasCancelled()
	})
}

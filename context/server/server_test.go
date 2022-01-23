package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store was cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() { s.cancelled = true }

type SpyResponseWriter struct{ written bool }

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) { s.written = true }

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}

		request := httptest.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()

		Server(store).ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q want %q", response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work on request cancel", func(t *testing.T) {
		data := "hello world"
		store := &SpyStore{response: data}

		request := httptest.NewRequest("GET", "/", nil)
		requestContext, cancel := context.WithCancel(request.Context())
		request = request.WithContext(requestContext)
		response := &SpyResponseWriter{}

		time.AfterFunc(5*time.Millisecond, cancel)
		Server(store).ServeHTTP(response, request)

		if response.written {
			t.Errorf("received a response")
		}
	})
}

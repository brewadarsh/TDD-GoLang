package Context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Store for testing purpose only.
type SpyStore struct {
	response  string
	cancelled bool
}

// Conforming to Store interface.
func (store *SpyStore) Fetch() string {
	time.Sleep(time.Second * 1) // Sleep for 1 second.
	return store.response
}

func (store *SpyStore) Cancel() {
	store.cancelled = true
}

func TestContext(t *testing.T) {
	t.Run("Running without cancel", func(t *testing.T) {
		data := "hello, world!"
		server := Serve(&SpyStore{response: data})
		// Create dummy response and request.
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf("Expected %v but got %v", data, res.Body.String())
		}
	})
	t.Run("Running with cancel", func(t *testing.T) {
		data := "hello, world!"
		store := SpyStore{response: data}
		server := Serve(&store)

		// Create the dummy request.
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingContext, triggerCancel := context.WithCancel(request.Context())

		// Trigger cancel before the store fetches.
		time.AfterFunc(time.Millisecond*200, triggerCancel)

		// Update the context of the request.
		request = request.WithContext(cancellingContext)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		if !store.cancelled {
			t.Errorf("Request not cancelled")
		}
	})
}

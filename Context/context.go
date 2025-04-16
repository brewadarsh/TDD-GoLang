package Context

import (
	"fmt"
	"net/http"
)

// The store interface.
type Store interface {
	Fetch() string
	Cancel()
}

// Setup the server.
func Serve(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channel := make(chan string)
		requestContext := r.Context()

		// Create the go-routine to fetch the data.
		go func() {
			channel <- store.Fetch()
		}()

		// Select from fetch or cancel, whichever finishes first.
		select {
		case storeResponse := <-channel:
			fmt.Fprint(w, storeResponse)
		case <-requestContext.Done():
			store.Cancel()
		}
	}
}

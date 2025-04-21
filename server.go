package main

import (
	"fmt"
	"net/http"
	"strings"
)

// The store interface.
type Store interface {
	// Get the score of the house.
	GetScore(house string) string
	// Store the update calls.
	RecordUpdate(house string)
}

// The server.
type Server struct {
	store Store
}

// The request handler for [Server].
func (server *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	house := strings.TrimPrefix(request.URL.Path, "/houses/")
	switch request.Method {
	case http.MethodPost:
		server.update(house, writer)
	default:
		server.read(house, writer)
	}
}

// Update the score.
func (server *Server) update(house string, writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusAccepted)
	server.store.RecordUpdate(house)
}

// Return the score.
func (server *Server) read(house string, writer http.ResponseWriter) {
	score := server.store.GetScore(house)
	if score == "" {
		writer.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(writer, score)
}

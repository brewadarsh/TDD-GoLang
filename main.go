package main

import (
	"log"
	"net/http"
)

// The in-memory storage.
type MemoryStore struct {
	scores map[string]string
}

// Get the score according to the house.
func (store *MemoryStore) GetScore(house string) string {
	return store.scores[house]
}

// Store the update calls made to an house.
func (store *MemoryStore) RecordUpdate(house string) {}

func main() {
	store := &MemoryStore{}
	server := &Server{store: store}
	log.Fatal(http.ListenAndServe(":3000", server))
}

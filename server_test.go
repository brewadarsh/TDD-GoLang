package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	scores  map[string]string
	updates []string
}

func (store *StubStore) GetScore(house string) string {
	s := store.scores[house]
	return s
}

func (store *StubStore) RecordUpdate(house string) {
	store.updates = append(store.updates, house)
}

func TestServer(t *testing.T) {
	store := StubStore{
		scores: map[string]string{
			"Gryffindor": "130",
			"Slytherin":  "120",
		},
	}
	server := &Server{&store}

	t.Run("returns the score of Gryffindor", func(t *testing.T) {
		house := "Gryffindor"
		request, _ := createGETRequest(house)
		responseRecorder := httptest.NewRecorder()
		server.ServeHTTP(responseRecorder, request)
		score := responseRecorder.Body.String()
		assertScore(t, score, "130")
	})
	t.Run("returns the score of Slytherin", func(t *testing.T) {
		house := "Slytherin"
		request, _ := createGETRequest(house)
		responseRecorder := httptest.NewRecorder()
		server.ServeHTTP(responseRecorder, request)
		score := responseRecorder.Body.String()
		assertScore(t, score, "120")
	})
	t.Run("returns 404 if the house is not found in the store", func(t *testing.T) {
		house := "Hufflepuff"
		request, _ := createGETRequest(house)
		responseRecorder := httptest.NewRecorder()
		server.ServeHTTP(responseRecorder, request)
		assertStatus(t, responseRecorder.Code, http.StatusNotFound)
	})
	t.Run("returns 202 if the post request is accepted", func(t *testing.T) {
		house := "Slytherin"
		request, _ := createPOSTRequest(house)
		responseRecorder := httptest.NewRecorder()
		server.ServeHTTP(responseRecorder, request)
		assertStatus(t, responseRecorder.Code, http.StatusAccepted)
	})
}

func TestServerStore(t *testing.T) {
	store := StubStore{}
	server := &Server{store: &store}

	t.Run("store the update calls", func(t *testing.T) {
		house := "Slytherin"
		request, _ := createPOSTRequest(house)
		recorder := httptest.NewRecorder()

		server.ServeHTTP(recorder, request)

		assertStatus(t, recorder.Code, http.StatusAccepted)
		if len(store.updates) != 1 {
			t.Errorf("expected update calls to be 1 but got %d", len(store.updates))
		}
		if store.updates[0] != house {
			t.Errorf("expected to record update for house %q not %q", house, store.updates[0])
		}
	})
}

// Create the GET request.
func createGETRequest(house string) (*http.Request, error) {
	endpoint := fmt.Sprintf("/houses/%s", house)
	return http.NewRequest(http.MethodGet, endpoint, nil)
}

// Create the POST request.
func createPOSTRequest(house string) (*http.Request, error) {
	endpoint := fmt.Sprintf("/houses/%s", house)
	return http.NewRequest(http.MethodPost, endpoint, nil)
}

// Asserts expected and actual scores.
func assertScore(t *testing.T, expectedScore, actualScore string) {
	t.Helper()
	if expectedScore != actualScore {
		t.Errorf("Expected %q but got %q", expectedScore, actualScore)
	}
}

// Asserts expected and actual statuses.
func assertStatus(t *testing.T, expectedStatus, actualStatus int) {
	t.Helper()
	if expectedStatus != actualStatus {
		t.Errorf("got status %d want %d", expectedStatus, actualStatus)
	}
}

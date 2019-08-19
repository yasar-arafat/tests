package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	winCalls []string
	score    map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.score[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		nil,
		map[string]int{
			"Yasar":  20,
			"Arbaaz": 10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("returns Yasar's score", func(t *testing.T) {
		request := newGetRequest("Yasar")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)

	})

	t.Run("returns Arbaaz's score", func(t *testing.T) {
		request := newGetRequest("Arbaaz")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("return 404 on missing player", func(t *testing.T) {
		request := newGetRequest("Suzan")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)

	})

}

func TestStoreWins(t *testing.T) {

	player := "Nasir"
	store := StubPlayerStore{
		nil,
		map[string]int{},
	}

	server := &PlayerServer{&store}

	t.Run("it returns accepted a POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/"+player, nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], player)
		}

	})

}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {

	server := PlayerServer{NewInMemoryPlayerStore()}
	player := "Yasar"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetRequest(player))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")

}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct response, got %d, want %d", got, want)
	}
}

func newGetRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return request
}
func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got '%v', want '%v'", got, want)
	}
}

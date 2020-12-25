package poker

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

// StubPlayerStore testing storing scores information about players
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

// GetPlayerScore testing returns the player's score from the file system store
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

// RecordWin testing saves the player win to file
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// GetLeague testing returns the league from the static store
func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// AssertPlayerWin asserts if the player won the game
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}

// AssertResponseBody asserts if the body has correct response
func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wan %q", got, want)
	}
}

// AssertStatus asserts if the status has correct response
func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

// AssertContentType asserts if the resquest has correct content-type
func AssertContentType(t *testing.T, r *httptest.ResponseRecorder, want string) {
	t.Helper()
	if r.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, r.Result().Header)
	}
}

// CreateTempFile creates a temporary file
func CreateTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp fle %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

// AssertNoError asserts if there is no error
func AssertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

// AssertLeague asserts if the league is correct
func AssertLeague(t *testing.T, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// AssertScoreEquals asserts if the score is correct
func AssertScoreEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("the scores are different, got %v, want %v", got, want)
	}
}

// GetLeagueFromResponse return the league from the response
func GetLeagueFromResponse(t *testing.T, body io.Reader) (league League) {
	t.Helper()
	league, err := NewLeague(body)

	if err != nil {
		t.Fatalf("Unable to parse the response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

// NewLeagueRequest creates a new request for league endpoint
func NewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

// NewPostWinRequest creates a new request for register player wins
func NewPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// NewGetScoreRequest creates a new request for getting player score
func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

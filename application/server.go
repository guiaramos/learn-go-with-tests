package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// PlayerStore stores scores information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// Player is a collection of properties for a player
type Player struct {
	Name string
	Wins int
}

// FileSystemPlayerStore is a collection of properties and methods for storing a player
type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// NewPlayerServer creates a new server for player store
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router

	return p
}

// GetLeague returns the league from the static store
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

// GetPlayerScore returns the player's score from the file system store
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

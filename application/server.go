package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore stores scores information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
}

// ServeHTTP defines the server
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

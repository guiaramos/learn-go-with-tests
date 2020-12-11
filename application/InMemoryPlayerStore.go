package main

// InMemoryPlayerStore temporally stores information about players
type InMemoryPlayerStore struct {
	store map[string]int
}

// NewInMemoryPlayerStore create a new in memory store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// GetPlayerScore temporally return player score from store
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin temporally add player score to store
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

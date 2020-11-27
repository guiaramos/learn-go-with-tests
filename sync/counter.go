package sync

import "sync"

// Counter creates a collection for counting
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter creates a new counter collection
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the counter collection
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value return the value of the counter collection
func (c *Counter) Value() int {
	return c.value
}

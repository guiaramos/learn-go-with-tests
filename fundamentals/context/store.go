package context

import (
	"context"
	"fmt"
	"net/http"
)

// Store is a named collection for store methods
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server return a http server based on store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error
		}

		fmt.Fprint(w, data)
	}
}

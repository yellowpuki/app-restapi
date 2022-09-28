package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
	"github.com/yellowpuki/app-restapi/internal/app/store"
)

// Store ...
type Store struct {
	db             *sql.DB
	userReposytory *UserReposytory
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserReposytory {
	if s.userReposytory != nil {
		return s.userReposytory
	}

	s.userReposytory = &UserReposytory{
		store: s,
	}

	return s.userReposytory
}

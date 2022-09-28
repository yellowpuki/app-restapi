package teststore

import (
	"github.com/yellowpuki/app-restapi/internal/app/model"
	"github.com/yellowpuki/app-restapi/internal/app/store"
)

// Store ...
type Store struct {
	userReposytory *UserReposytory
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserReposytory {
	if s.userReposytory != nil {
		return s.userReposytory
	}

	s.userReposytory = &UserReposytory{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userReposytory
}

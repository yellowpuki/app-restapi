package store

import "github.com/yellowpuki/app-restapi/internal/app/model"

// UserReposytory ...
type UserReposytory interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

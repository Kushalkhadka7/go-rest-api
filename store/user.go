package store

import (
	"go-rest-api/db"
	error "go-rest-api/errors"
	"go-rest-api/modal"
)

// UserStore new user Modal
type UserStore struct {
	*db.Conn
}

// NewUser creates new modal for user.
func NewUser(db *db.Conn) *UserStore {
	return &UserStore{
		db,
	}
}

// Get all the list of users.
func (store *UserStore) Get() ([]modal.User, *error.Response) {
	var user []modal.User

	if err := store.Find(&user).Error; err != nil {
		return nil, error.NotFound(err)
	}

	return user, nil
}

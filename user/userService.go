package user

import (
	error "go-rest-api/errors"
	"go-rest-api/modal"

	"github.com/sirupsen/logrus"
)

// Service handles computations for users.
type Service interface {
	Get() ([]modal.User, *error.Response)
}

// ServiceStore connects modal and service.
type ServiceStore struct {
	Store Service
}

// NewService creates new user service.
func NewService(store Service, log *logrus.Logger) *ServiceStore {

	return &ServiceStore{
		Store: store,
	}
}

// Get list of all users.
func (us *ServiceStore) Get() ([]modal.User, *error.Response) {

	user, err := us.Store.Get()
	if err != nil {
		return nil, err
	}

	return user, nil
}

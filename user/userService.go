package user

import error "go-rest-api/errors"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (us *Service) get() (map[string]string, *error.Response) {
	user := make(map[string]string)
	user["fname"] = "kushal"

	return user, nil
}

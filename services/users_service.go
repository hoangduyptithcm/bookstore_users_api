package services

import (
	"net/http"

	"github.com/hoangduyptithcm/bookstore_users_api/domain/users"
	"github.com/hoangduyptithcm/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return nil, nil

	return &user, nil

	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}

package user

import (
	errors "yu-croco/ddd_on_golang/pkg/errors"
)

type UserRepository interface {
	FindById(id UserId) (*User, *errors.AppError)
}

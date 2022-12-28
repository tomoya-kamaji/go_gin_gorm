package group

import (
	errors "yu-croco/ddd_on_golang/pkg/errors"
)

type GroupRepository interface {
	// FindById(id GroupId) (*Group, *errors.AppError)
	Save(user *Group) (*Group, *errors.AppError)
}

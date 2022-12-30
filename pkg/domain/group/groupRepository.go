package group

import (
	errors "yu-croco/ddd_on_golang/pkg/errors"
)

type GroupRepository interface {
	Save(user *Group) (*Group, *errors.AppError)
}

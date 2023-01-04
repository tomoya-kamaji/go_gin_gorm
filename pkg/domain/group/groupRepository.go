package group

import (
	errors "yu-croco/ddd_on_golang/pkg/errors"
)

type GroupRepository interface {
	Save(group *Group) (*Group, *errors.AppError)
	FindById(id GroupId) *Group
}

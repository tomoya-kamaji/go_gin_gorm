package query

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
)

type UserQuery interface {
	FindAll() *model2.Users
}

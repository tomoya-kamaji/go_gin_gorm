package query

import (
	"yu-croco/ddd_on_golang/pkg/domain/user"
)

type UserQuery interface {
	FindAll() *user.User
}

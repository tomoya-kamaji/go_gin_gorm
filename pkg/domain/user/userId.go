package user

import (
	"yu-croco/ddd_on_golang/pkg/errors"
	"yu-croco/ddd_on_golang/pkg/lib/uid"
)

type UserId int

func NewUserId(id int) (*UserId, *errors.AppError) {
	userId := UserId(id)
	return &userId, nil
}

func CreateUserId() *UserId {
	userId := UserId(uid.CreateUid())
	return &userId
}

package user

import (
	"math/rand"
	"yu-croco/ddd_on_golang/pkg/errors"
)

type UserId int

func NewUserId(id int) (*UserId, *errors.AppError) {
	userId := UserId(id)
	return &userId, nil
}

func CreateUserId() *UserId {
	userId := UserId(rand.Intn(10000))
	return &userId
}

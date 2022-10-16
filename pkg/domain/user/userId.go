package user

import (
	"yu-croco/ddd_on_golang/pkg/errors"
)
type UserId int

func NewUserId(id int) (*UserId, *errors.AppError) {
	if id <= 0 {
		err := errors.NewAppError("UserIdは1以上の値にしてください")
		return nil, &err
	}
	userId := UserId(id)
	return &userId, nil
}

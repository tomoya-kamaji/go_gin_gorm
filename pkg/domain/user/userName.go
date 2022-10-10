package user

import (
	"yu-croco/ddd_on_golang/pkg/errors"
)
type UserName string

func NewUserName(name string) (*UserName, *errors.AppError) {
	if len(name) <= 15  {
		err := errors.NewAppError("UseNameは15文字以下にしてください")
		return nil, &err
	}
	userName := UserName(name)
	return &userName, nil
}

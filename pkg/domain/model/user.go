package model

import (
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
)

type User struct {
	Id              UserId           `json:"userId"`
	Name            UserName         `json:"userName"`
}

// 完全コンストラクタのための初期化処理サンプル
func NewUserId(id int) (*UserId, *errors2.AppError) {
	if id <= 0 {
		err := errors2.NewAppError("UserIdは1以上の値にしてください")
		return nil, &err
	}

	userId := UserId(id)
	return &userId, nil
}

type UserId int
type UserName string


type Users []User
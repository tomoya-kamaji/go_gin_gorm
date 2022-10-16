package dto

import (
	"yu-croco/ddd_on_golang/pkg/domain/user"
)

type UserEntity struct {
	ID           uint `json:"id" binding:"required"`
	Name         string
}

func (u *UserEntity) TableName() string {
	return "user"
}

func (u *UserEntity) ConvertToModel() *user.User {
	return user.NewUser(user.UserId(u.ID), user.UserName(u.Name))
}
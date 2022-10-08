package dto

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
)

type User struct {
	ID           uint `json:"id" binding:"required"`
	Name         string
}

type Users []User

func (u *User) ConvertToModel() *model2.User {
	return &model2.User{
		Id:           model2.UserId(u.ID),
		Name:         model2.UserName(u.Name),
	}
}

func (users Users) ConvertToModel() *model2.Users {
	result := make(model2.Users, len(users))

	for idx, user := range users {
		userModel := model2.User{
			Id:              model2.UserId(user.ID),
			Name:            model2.UserName(user.Name),
		}
		result[idx] = userModel
	}

	return &result
}

package dto

import "yu-croco/ddd_on_golang/pkg/domain/user"

type GroupUsersEntity struct {
	GroupId uint
	UserId  uint
}

func (g *GroupUsersEntity) TableName() string {
	return "group_users"
}

func (u *UserEntity) ConvertToModel() *user.User {
	return user.NewUser(user.UserId(u.ID), user.UserName(u.Name))
}

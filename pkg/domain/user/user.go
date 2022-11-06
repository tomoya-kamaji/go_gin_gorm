package user

import "yu-croco/ddd_on_golang/pkg/errors"

type User struct {
	Id   UserId   `json:"userId"`
	Name UserName `json:"userName"`
}

func NewUser(id UserId, name UserName) *User {
	return &User{Id: id, Name: name}
}

type Users []User

func (user *User) ChangeName(name string) (*User, *errors.AppError) {
	userName, err := NewUserName(name)
	if err.HasErrors() {
		return nil, err
	}
	user.Name = *userName
	return user, nil
}

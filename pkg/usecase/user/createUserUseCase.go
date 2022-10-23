package usecase

import (
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/errors"
)

type CreateUserUsecaseImpl struct {
	UserName          string
	UserRepository  user.UserRepository
}

type CreateUserUsecase interface {
	Run() (*user.User, *errors.AppError)
}

func NewCreateUserUsecaseImpl(userName string, userRepository user.UserRepository) CreateUserUsecase {
	return CreateUserUsecaseImpl{
		UserName:          userName,
		UserRepository:  userRepository,
	}
}

func (impl CreateUserUsecaseImpl) Run() (*user.User, *errors.AppError) {
	userId := user.CreateUserId()
	userName,userNameErr := user.NewUserName(impl.UserName)
	if userNameErr.HasErrors() {
		return nil, userNameErr
	}
	user := user.NewUser(*userId,*userName)
	saveUser, userSaveErr := impl.UserRepository.Save(user)
	if userSaveErr.HasErrors() {
		return nil, userSaveErr
	}
	return saveUser, nil
}

package usecase

import (
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/errors"
)

type createUserUsecaseImpl struct {
	UserName          string
	UserRepository  user.UserRepository
}

type createUserUsecase interface {
	Run() (*user.User, *errors.AppError)
}

func NewCreateUserUsecaseImpl(userName string, userRepository user.UserRepository) createUserUsecase {
	return createUserUsecaseImpl{
		UserName:          userName,
		UserRepository:  userRepository,
	}
}

func (impl createUserUsecaseImpl) Run() (*user.User, *errors.AppError) {
	userId, err := user.CreateUserId()
	if(err.HasErrors()){
		return nil, err
	}
	user := user.NewUser(*userId,user.UserName(impl.UserName))
	saveUser, userSaveErr := impl.UserRepository.Save(user)
	if userSaveErr.HasErrors() {
		return nil, userSaveErr
	}
	return saveUser, nil
}

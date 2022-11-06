package usecase

import (
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/errors"
)

type UpdateUserUsecaseImpl struct {
	UserId         user.UserId
	UserName       string
	UserRepository user.UserRepository
}

type UpdateUserUsecase interface {
	Run() (*user.User, *errors.AppError)
}

func NewUpdateUserUsecaseImpl(userId user.UserId, userName string, userRepository user.UserRepository) UpdateUserUsecase {
	return UpdateUserUsecaseImpl{
		UserId:         userId,
		UserName:       userName,
		UserRepository: userRepository,
	}
}

func (impl UpdateUserUsecaseImpl) Run() (*user.User, *errors.AppError) {
	user, findErr := impl.UserRepository.FindById(impl.UserId)
	if findErr.HasErrors() {
		return nil, findErr
	}

	user.ChangeName(impl.UserName)
	saveUser, userSaveErr := impl.UserRepository.Save(user)
	if userSaveErr.HasErrors() {
		return nil, userSaveErr
	}
	return saveUser, nil
}

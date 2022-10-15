package userUsecase

import (
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/errors"
)

type fetchUserDetailUsecaseImpl struct {
	UserId          user.UserId
	UserRepository  user.UserRepository
}

type fetchUserDetailUsecase interface {
	Run() (*user.User, *errors.AppError)
}

func NewFetchUserDetailUsecaseImpl(userId user.UserId, userRepository user.UserRepository) fetchUserDetailUsecase {
	return fetchUserDetailUsecaseImpl{
		UserId:          userId,
		UserRepository:  userRepository,
	}
}


func (impl fetchUserDetailUsecaseImpl) Run() (*user.User, *errors.AppError) {
	user, userFindErr := impl.UserRepository.FindById(impl.UserId)
	if userFindErr.HasErrors() {
		return nil, userFindErr
	}
	return user, nil
}

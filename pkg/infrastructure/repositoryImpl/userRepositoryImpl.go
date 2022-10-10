package repositoryImpl

import (
	"yu-croco/ddd_on_golang/pkg/domain/user"
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"
	"yu-croco/ddd_on_golang/pkg/infrastructure/dto"
)

type userRepositoryImpl struct{}

func NewUserRepositoryImpl() user.UserRepository {
	return &userRepositoryImpl{}
}

func (repositoryImpl *userRepositoryImpl) FindById(id user.UserId) (*user.User, *errors2.AppError) {
	db := infrastructure2.GetDB()
	var err errors2.AppError
	userEntity := dto.UserEntity{}

	if db.Find(&userEntity, dto.UserEntity{ID: uint(id)}).RecordNotFound() {
		err = notFoundUserError(id)
		return nil, &err
	}

	return userEntity.ConvertToModel(), nil
}


func notFoundUserError(id user.UserId) errors2.AppError {
	return errors2.NewAppError("userは見つかりませんでした")
}

package repositoryImpl

import (
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/errors"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"
	"yu-croco/ddd_on_golang/pkg/infrastructure/dto"
)

type userRepositoryImpl struct{}

func NewUserRepositoryImpl() user.UserRepository {
	return &userRepositoryImpl{}
}

func (repositoryImpl *userRepositoryImpl) FindById(id user.UserId) (*user.User, *errors.AppError) {
	db := infrastructure2.GetDB()
	var err errors.AppError
	userEntity := dto.UserEntity{}

	if db.Find(&userEntity, dto.UserEntity{ID: uint(id)}).RecordNotFound() {
		err = notFoundUserError(id)
		return nil, &err
	}

	return userEntity.ConvertToModel(), nil
}

func (repositoryImpl *userRepositoryImpl) Save(user *user.User) (*user.User, *errors.AppError) {
	db := infrastructure2.GetDB()
	userEntity := dto.UserEntity{}

	if db.Find(&userEntity, dto.UserEntity{ID: uint(user.Id)}).RecordNotFound() {
		userEntity.ID = uint(user.Id)
		userEntity.Name = string(user.Name)
	} else {
		userEntity.Name = string(user.Name)
	}
	db.Save(&userEntity)
	return userEntity.ConvertToModel(), nil
}


func notFoundUserError(id user.UserId) errors.AppError {
	return errors.NewAppError("userは見つかりませんでした")
}

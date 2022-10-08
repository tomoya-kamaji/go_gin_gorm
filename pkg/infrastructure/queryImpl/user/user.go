package user

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"
	dto2 "yu-croco/ddd_on_golang/pkg/infrastructure/dto"
	query2 "yu-croco/ddd_on_golang/pkg/query"
)

type userQueryImpl struct{}

func NewUserQueryImpl() query2.UserQuery {
	return &userQueryImpl{}
}

func (repo userQueryImpl) FindAll() *model2.Users {
	db := infrastructure2.GetDB()
	userDaos := dto2.Users{}

	db.Preload("Users").Find(&userDaos)


	return userDaos.ConvertToModel()
}

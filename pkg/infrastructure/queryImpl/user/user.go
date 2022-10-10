package user

import (
 "yu-croco/ddd_on_golang/pkg/domain/user"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"
	"yu-croco/ddd_on_golang/pkg/infrastructure/dto"
	query2 "yu-croco/ddd_on_golang/pkg/query"
)

type userQueryImpl struct{}

func NewUserQueryImpl() query2.UserQuery {
	return &userQueryImpl{}
}

func (repo userQueryImpl) FindAll() *user.User {
	db := infrastructure2.GetDB()
	userEntity := dto.UserEntity{}

	db.Find(&userEntity)

	return userEntity.ConvertToModel()
}

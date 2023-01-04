package repositoryImpl

import (
	"fmt"
	"yu-croco/ddd_on_golang/pkg/domain/group"
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/errors"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"
	"yu-croco/ddd_on_golang/pkg/infrastructure/dto"
)

type groupRepositoryImpl struct{}

func NewGroupRepositoryImpl() group.GroupRepository {
	return &groupRepositoryImpl{}
}

func (repositoryImpl *groupRepositoryImpl) Save(group *group.Group) (*group.Group, *errors.AppError) {
	db := infrastructure2.GetTestDB()
	groupEntity := dto.GroupEntity{}
	userEntities := []dto.UserEntity{}

	db.Where("id IN (?)", group.UserIds).Find(&userEntities)
	groupEntity.ID = uint(group.Id)
	groupEntity.Name = group.Name
	groupEntity.Users = userEntities
	db.Save(&groupEntity)
	return ConvertToModel(&groupEntity), nil
}

func (repositoryImpl *groupRepositoryImpl) FindById(id group.GroupId) *group.Group {
	db := infrastructure2.GetTestDB()
	groupEntity := dto.GroupEntity{}

	db.Preload("Users").Find(&groupEntity, dto.GroupEntity{ID: uint(id)})
	fmt.Printf("groupEntity: %v\n", groupEntity)
	return ConvertToModel(&groupEntity)
}

func ConvertToModel(g *dto.GroupEntity) *group.Group {
	return group.Reconstruct(group.GroupId(g.ID), g.Name, []user.UserId{})
}

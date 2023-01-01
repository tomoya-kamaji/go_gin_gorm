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
	db := infrastructure2.GetDB()
	groupEntity := dto.GroupEntity{}

	db.Debug().Preload("Users").Find(&groupEntity)
	fmt.Printf("groupEntity: %v\n", groupEntity)
	fmt.Printf("groupEntity: %v\n", &groupEntity)

	groupEntity.ID = uint(group.Id)
	groupEntity.Name = group.Name
	db.Debug().Save(&groupEntity)
	return ConvertToModel(&groupEntity), nil
}

func ConvertToModel(g *dto.GroupEntity) *group.Group {
	return group.Reconstruct(group.GroupId(g.ID), g.Name, []user.UserId{})
}

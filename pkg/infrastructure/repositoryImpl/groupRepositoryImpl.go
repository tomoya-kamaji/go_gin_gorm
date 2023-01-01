package repositoryImpl

import (
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
	groupUsersEntity := dto.GroupUsersEntity{}

	if db.Find(&groupEntity, dto.GroupEntity{ID: uint(group.Id)}).RecordNotFound() {
		groupEntity.ID = uint(group.Id)
		groupEntity.Name = string(group.Name)
	} else {
		groupEntity.Name = string(group.Name)
	}

	// db.Where("group_id = ?", "%1%").Delete(groupUsersEntity)
	for _, user := range group.UserIds {
		groupUsersEntity.GroupId = uint(group.Id)
		groupUsersEntity.UserId = uint(user)
		db.Save(&groupUsersEntity)
	}

	db.Save(&groupEntity)
	return ConvertToModel(&groupEntity), nil
}

func ConvertToModel(g *dto.GroupEntity) *group.Group {
	return group.Reconstruct(group.GroupId(g.ID), g.Name, []user.UserId{})
}

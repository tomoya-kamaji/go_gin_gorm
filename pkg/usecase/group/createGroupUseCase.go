package usecase

import (
	"yu-croco/ddd_on_golang/pkg/domain/group"
	"yu-croco/ddd_on_golang/pkg/errors"
)

type createGroupUseCaseImpl struct {
	GroupName       string
	GroupRepository group.GroupRepository
}

type createGroupUseCase interface {
	Run() (*group.Group, *errors.AppError)
}

func NewCreateGroupUseCaseImpl(groupName string, groupRepository group.GroupRepository) createGroupUseCase {
	return createGroupUseCaseImpl{
		GroupName:       groupName,
		GroupRepository: groupRepository,
	}
}

func (impl createGroupUseCaseImpl) Run() (*group.Group, *errors.AppError) {
	group := group.CreateGroup(impl.GroupName)
	saveGroup, groupSaveErr := impl.GroupRepository.Save(group)
	if groupSaveErr.HasErrors() {
		return nil, groupSaveErr
	}
	return saveGroup, nil
}

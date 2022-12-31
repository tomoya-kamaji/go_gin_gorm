package group

import (
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/errors"
)

type Group struct {
	Id      GroupId       `json:"groupId"`
	Name    string        `json:"groupName"`
	UserIds []user.UserId `json:"userIds"`
}

func NewGroup(id GroupId, name string) *Group {
	return &Group{Id: id, Name: name, UserIds: make([]user.UserId, 0)}
}

func CreateGroup(name string) *Group {
	id := CreateGroupId()
	return &Group{Id: *id, Name: name, UserIds: make([]user.UserId, 0)}
}

func (group *Group) AddUser(userId user.UserId) *errors.AppError {
	//userIdsにuserIdが存在するか確認
	for _, id := range group.UserIds {
		if id == userId {
			err := errors.NewAppError("同一ユーザは追加できない")
			return &err
		}
	}
	group.UserIds = append(group.UserIds, userId)
	return nil
}

func (group *Group) RemoveUser(userId user.UserId) {
	for i, id := range group.UserIds {
		if id == userId {
			group.UserIds = append(group.UserIds[:i], group.UserIds[i+1:]...)
			break
		}
	}
}

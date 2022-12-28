package group

import "yu-croco/ddd_on_golang/pkg/domain/user"

type Group struct {
	Id      GroupId       `json:"groupId"`
	Name    string        `json:"groupName"`
	UserIds []user.UserId `json:"userIds"`
}

func NewGroup(id GroupId, name string) *Group {
	return &Group{Id: id, Name: name, UserIds: make([]user.UserId, 0)}
}

func (group *Group) AddUser(userId user.UserId) {
	group.UserIds = append(group.UserIds, userId)
}

package group

import (
	"yu-croco/ddd_on_golang/pkg/errors"
	"yu-croco/ddd_on_golang/pkg/lib/uid"
)

type GroupId int

func NewGroupId(id int) (*GroupId, *errors.AppError) {
	groupId := GroupId(id)
	return &groupId, nil
}

func CreateGroupId() *GroupId {
	groupId := GroupId(uid.CreateUid())
	return &groupId
}

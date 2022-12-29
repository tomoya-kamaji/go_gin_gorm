package group

import (
	"testing"
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/errors"

	"github.com/google/go-cmp/cmp"
)

func TestNewGroup(t *testing.T) {
	cases := map[string]struct {
		Id      GroupId
		Name    string
		want    *Group
		wantErr error
	}{
		"NewGroup": {
			Id:   GroupId(1),
			Name: "エンジニアチーム",
			want: &Group{Id: GroupId(1), Name: "エンジニアチーム", UserIds: make([]user.UserId, 0)},
		}}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res := NewGroup(tc.Id, tc.Name)
			if diff := cmp.Diff(res, tc.want); diff != "" {
				t.Errorf("X value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	cases := map[string]struct {
		group   *Group
		userId  user.UserId
		want    []user.UserId
		wantErr *errors.AppError
	}{
		"正常パターン：指定したユーザが追加される": {
			group:  NewGroup(GroupId(1), "エンジニアチーム"),
			userId: user.UserId(1),
			want:   make([]user.UserId, 0),
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res := tc.group.UserIds
			if diff := cmp.Diff(res, tc.want); diff != "" {
				t.Errorf("value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}

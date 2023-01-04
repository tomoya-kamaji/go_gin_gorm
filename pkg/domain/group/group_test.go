package group

import (
	"testing"
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/errors"

	"github.com/google/go-cmp/cmp"
)

func TestAddUser(t *testing.T) {
	cases := map[string]struct {
		group   *Group
		userId  user.UserId
		want    []user.UserId
		wantErr *errors.AppError
	}{
		"正常パターン：指定したユーザが追加される": {
			group:  NewGroup(GroupId(1), "エンジニアチーム", make([]user.UserId, 0)),
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

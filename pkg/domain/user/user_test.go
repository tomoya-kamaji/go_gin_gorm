package user

import (
	"testing"
	"yu-croco/ddd_on_golang/pkg/errors"
	"yu-croco/ddd_on_golang/pkg/lib/errorCheck"
	"yu-croco/ddd_on_golang/pkg/lib/pointerLib"

	"github.com/google/go-cmp/cmp"
)

func TestNewUser(t *testing.T) {
	cases := map[string]struct {
		Id      UserId
		Name    UserName
		want    *User
		wantErr error
	}{
		"NewUser": {
			Id:   UserId(1),
			Name: UserName("tomoya"),
			want: &User{Id: UserId(1), Name: UserName("tomoya")},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res := NewUser(tc.Id, tc.Name)

			if diff := cmp.Diff(res, tc.want); diff != "" {
				t.Errorf("X value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}

func TestNewUserName(t *testing.T) {
	cases := map[string]struct {
		name    string
		want    *UserName
		wantErr *errors.AppError
	}{
		"正常パターン": {
			name: "tomoya",
			want: pointerLib.ToPointer(UserName("tomoya")),
		},
		"15文字以上の場合エラー": {
			name:    "1234567890123456",
			wantErr: pointerLib.ToPointer(errors.NewAppError("UseNameは15文字以下にしてください")),
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := NewUserName(tc.name)
			if diff := cmp.Diff(res, tc.want); diff != "" {
				t.Errorf("value is mismatch (-want +got):%s\n", diff)
			}
			errorCheck.Check(t, err, tc.wantErr)
		})
	}
}

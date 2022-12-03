package parallelprocessing

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}



type User struct {
	Id   string
	Name string
}
type User2 struct {
	Id   string
	Name string
	Name2 string
}
func TestParallelprocessing(t *testing.T) {
	cases := map[string]struct {
		Name    string
		want    string
	}{
		"正常系": {
			Name: "釜地智也",
			want: "釜地智也",
		},
	}


	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			user1 := User{Id:"1",Name: "tomoya1"}
			user2 := User{Id:"2",Name: "tomoya2"}
			user3 := User{Id:"3",Name: "tomoya3"}
			users := []User{user1,user2,user3}

			q := make(chan User2)
			for _,user := range users {
				go func(user User) {
					q <- User2{Id:user.Id,Name:user.Name,Name2:user.Name}
				}(user)
			}
			a := <-q // receive from c

			fmt.Printf("a: %v\n", a)

			res := "hoge"
			if diff := cmp.Diff(res, tc.want); diff != "" {
				t.Errorf("X value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}

func convertUser2(user User) User2{
	return User2{Id:user.Id,Name:user.Name,Name2:user.Name}
}

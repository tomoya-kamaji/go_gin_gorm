package repositoryImpl

import (
	"fmt"
	"os"
	"testing"
	"yu-croco/ddd_on_golang/pkg/domain/group"
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/infrastructure"
)

func TestGroupRepositoryImplSave(t *testing.T) {
	groupRepositoryImpl := NewGroupRepositoryImpl()
	userRepositoryImpl := NewUserRepositoryImpl()
	testdb := infrastructure.GetTestDB()

	// save
	t.Run("新規で1レコード作成される", func(t *testing.T) {
		group := group.CreateGroup("テストグループ")

		groupRepositoryImpl.Save(group)
		fmt.Printf("userRepositoryImpl: %v\n", userRepositoryImpl)
		// nameが一致しているレコードがあるか
		if testdb.Debug().Where("name = ?", group.Name).First(&group).RecordNotFound() {
			t.Errorf("レコードが作成されていない")
		}
	})
	// get
	t.Run("groupsを取得する", func(t *testing.T) {
		groupRepositoryImpl := NewGroupRepositoryImpl()
		userRepositoryImpl := NewUserRepositoryImpl()

		user1 := user.CreateUser("ともや")
		userRepositoryImpl.Save(user1)
		user2 := user.CreateUser("さくら")
		userRepositoryImpl.Save(user2)

		group := group.CreateGroup("テストグループ")
		group.AddUser(user1.Id)
		group.AddUser(user2.Id)
		fmt.Printf("グループについて: %v\n", group)
		groupRepositoryImpl.Save(group)

		result := groupRepositoryImpl.FindById(group.Id)
		fmt.Printf("result: %v\n", result)
	})
}

func factory() {

}

func TestMain(m *testing.M) {
	resource, pool := infrastructure.CreateContainer()
	testdb := infrastructure.ConnectDB(pool)
	exitVal := m.Run()
	infrastructure.CloseContainer(resource, pool)
	testdb.Close()
	os.Exit(exitVal)
}

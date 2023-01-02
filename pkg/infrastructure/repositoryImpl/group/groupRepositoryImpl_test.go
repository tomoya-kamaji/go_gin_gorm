package repositoryImpl

import (
	"testing"
	"yu-croco/ddd_on_golang/pkg/domain/group"
	"yu-croco/ddd_on_golang/pkg/infrastructure"
)

func TestGroupRepositoryImplSave(t *testing.T) {
	repository := NewGroupRepositoryImpl()

	t.Run("新規で1レコード作成される", func(t *testing.T) {
		infrastructure.InitTest()

		group := group.CreateGroup("テストグループ")
		repository.Save(group)
		// result := repository.FindById(group.Id)
		// fmt.Printf("result: %v\n", result)
	})
}


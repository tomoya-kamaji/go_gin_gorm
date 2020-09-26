package monster

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/dto"
	"yu-croco/ddd_on_golang/app/query"
)

type MonsterQueryImpl struct{}

func NewMonsterQueryImpl() query.MonsterQuery {
	return &MonsterQueryImpl{}
}

func (repo MonsterQueryImpl) FindAll() *model.Monsters {
	db := infrastructure.GetDB()
	monsterDaos := dto.Monsters{}

	db.Preload("Materials").Find(&monsterDaos)

	return monsterDaos.ConvertToModel()
}

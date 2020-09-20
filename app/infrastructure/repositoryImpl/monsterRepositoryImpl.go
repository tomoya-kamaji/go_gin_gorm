package repositoryImpl

import (
	"github.com/jinzhu/gorm"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/infrastructure/dao"
)

func FindMonsterBy(db *gorm.DB, id int) *model.Monster {
	monsterDao := dao.Monster{}
	if db.First(&monsterDao, id).RecordNotFound() {
	}

	return monsterDao.ConvertToModel()
}

func UpdateMonster(db *gorm.DB, monster dao.Monster) *model.Monster {
	var monsterDao *dao.Monster

	if db.First(&monsterDao, int(monster.ID)).RecordNotFound() {
	}
	monsterDao.Life = monster.Life

	db.Save(&monsterDao)
	return monsterDao.ConvertToModel()
}

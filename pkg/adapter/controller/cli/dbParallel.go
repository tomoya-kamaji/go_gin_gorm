package controller

import (
	"fmt"
	"sync"
	"time"
	"yu-croco/ddd_on_golang/pkg/domain/user"
	"yu-croco/ddd_on_golang/pkg/infrastructure"
	"yu-croco/ddd_on_golang/pkg/infrastructure/dto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ClisController struct{}

func (ctrl ClisController) Run(c *gin.Context) {
	db := infrastructure.GetDB()
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 5; i++ {
		go createUser(&wg,db)
	}
	fmt.Printf("経過: %vms\n", time.Since(now).Milliseconds())
}

func createUser(
	wg *sync.WaitGroup,
	db *gorm.DB,
){
	defer wg.Done()
	userEntities := []dto.UserEntity{}
	for i := 0; i < 1000; i++ {
		userEntity := dto.UserEntity{}
		userID := user.CreateUserId()
		user := user.NewUser(*userID, user.UserName("tomoya"))
		if db.Find(&userEntity, dto.UserEntity{ID: uint(user.Id)}).RecordNotFound() {
			userEntity.ID = uint(user.Id)
			userEntity.Name = string(user.Name)
		} else {
			userEntity.Name = string(user.Name)
		}
		userEntities = append(userEntities, userEntity)
	}
	for _, userEntity := range userEntities {
		db.Save(&userEntity)
	}
}

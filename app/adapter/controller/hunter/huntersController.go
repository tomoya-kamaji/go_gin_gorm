package hunter

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

type Controller struct{}

func (ctrl Controller) Show(c *gin.Context) {
	db := infrastructure.GetDB()
	hunterId := c.Param("id")
	id, _ := strconv.Atoi(hunterId)
	dbResult := repositoryImpl.FindHunterById(db, id)

	c.JSON(200, dbResult)
}

func (ctrl Controller) Attack(c *gin.Context) {

}

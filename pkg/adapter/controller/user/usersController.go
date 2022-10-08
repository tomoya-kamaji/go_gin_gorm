package user

import (
	helpers2 "yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	"github.com/gin-gonic/gin"
)

type UsersController struct{}


func (ctrl UsersController) Index(c *gin.Context) {
	slice := [] string{"Golang", "Java"}
	helpers2.Response(c, slice, nil)
}

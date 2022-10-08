package hunter

import (
	helpers2 "yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	user2 "yu-croco/ddd_on_golang/pkg/infrastructure/queryImpl/user"
	"github.com/gin-gonic/gin"
)

type UsersController struct{}


func (ctrl UsersController) Index(c *gin.Context) {
	result := user2.NewUserQueryImpl().FindAll()
	helpers2.Response(c, result, nil)
}

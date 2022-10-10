package hunter

import (
	helpers2 "yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/pkg/domain/user"
	user2 "yu-croco/ddd_on_golang/pkg/infrastructure/queryImpl/user"
	"yu-croco/ddd_on_golang/pkg/infrastructure/repositoryImpl"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}


func (ctrl UsersController) Index(c *gin.Context) {
	result := user2.NewUserQueryImpl().FindAll()
	helpers2.Response(c, result, nil)
}

func (ctrl UsersController) Detail(c *gin.Context) {
		userId, err := user.NewUserId(helpers2.ConvertToInt(c.Param("id")))

		if err.HasErrors() {
			helpers2.Response(c, nil, err)
		} else {
			repo := repositoryImpl.NewUserRepositoryImpl()
			result, errs := repo.FindById(*userId)
			helpers2.Response(c, result, errs)
		}
}

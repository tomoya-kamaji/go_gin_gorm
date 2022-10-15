package hunter

import (
	helpers2 "yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	userModel "yu-croco/ddd_on_golang/pkg/domain/user"
	userQuery "yu-croco/ddd_on_golang/pkg/infrastructure/queryImpl/user"
	"yu-croco/ddd_on_golang/pkg/infrastructure/repositoryImpl"
	userUsecase "yu-croco/ddd_on_golang/pkg/usecase/user"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}


func (ctrl UsersController) Index(c *gin.Context) {
	result := userQuery.NewUserQueryImpl().FindAll()
	helpers2.Response(c, result, nil)
}

func (ctrl UsersController) Detail(c *gin.Context) {
		userId, err := userModel.NewUserId(helpers2.ConvertToInt(c.Param("id")))

		if err.HasErrors() {
			helpers2.Response(c, nil, err)
		} else {
			result, errs := userUsecase.NewFetchUserDetailUsecaseImpl(*userId, repositoryImpl.NewUserRepositoryImpl()).Run()
			helpers2.Response(c, result, errs)
		}
}

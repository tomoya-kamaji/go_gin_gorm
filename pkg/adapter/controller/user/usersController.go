package controller

import (
	"yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/pkg/domain/user"
	queryImpl "yu-croco/ddd_on_golang/pkg/infrastructure/queryImpl/user"
	"yu-croco/ddd_on_golang/pkg/infrastructure/repositoryImpl"
	usecase "yu-croco/ddd_on_golang/pkg/usecase/user"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}


func (ctrl UsersController) Index(c *gin.Context) {
	result := queryImpl.NewUserQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}

func (ctrl UsersController) Detail(c *gin.Context) {
		userId, err := user.NewUserId(helpers.ConvertToInt(c.Param("id")))
		if err.HasErrors() {
			helpers.Response(c, nil, err)
		} else {
			result, errs := usecase.NewFetchUserDetailUsecaseImpl(*userId, repositoryImpl.NewUserRepositoryImpl()).Run()
			helpers.Response(c, result, errs)
		}
}

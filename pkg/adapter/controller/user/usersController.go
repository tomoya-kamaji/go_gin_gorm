package controller

import (
	"net/http"
	"yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/pkg/domain/user"
	queryImpl "yu-croco/ddd_on_golang/pkg/infrastructure/queryImpl/user"
	"yu-croco/ddd_on_golang/pkg/infrastructure/repositoryImpl"
	usecase "yu-croco/ddd_on_golang/pkg/usecase/user"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

type CreateUserRequest struct {
	Name  string `json:"name"`
}


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

func (ctrl UsersController) Create(c *gin.Context) {
		var requestJson CreateUserRequest
		if err := c.ShouldBindJSON(&requestJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, errs := usecase.NewCreateUserUsecaseImpl(requestJson.Name, repositoryImpl.NewUserRepositoryImpl()).Run()
		helpers.Response(c, result, errs)
}

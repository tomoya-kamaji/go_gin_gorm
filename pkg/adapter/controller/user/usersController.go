package controller

import (
	"net/http"
	"yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/pkg/domain/user"
	elasticsearch "yu-croco/ddd_on_golang/pkg/infrastructure/elasticSearch"
	queryImpl "yu-croco/ddd_on_golang/pkg/infrastructure/queryImpl/user"
	"yu-croco/ddd_on_golang/pkg/infrastructure/repositoryImpl"
	usecase "yu-croco/ddd_on_golang/pkg/usecase/user"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

type CreateUserRequest struct {
	Name string `json:"name"`
}

func (ctrl UsersController) Index(c *gin.Context) {
	result := queryImpl.NewUserQueryImpl().FindAll()
	helpers.Response(c, result, nil)
}

func (ctrl UsersController) Search(c *gin.Context) {
	esAdapter := elasticsearch.NewElasticSearchAdapter()
	esAdapter.CreateIndex()
}

func (ctrl UsersController) Detail(c *gin.Context) {
	userID, err := user.NewUserId(helpers.ConvertToInt(c.Param("id")))
	if err.HasErrors() {
		helpers.Response(c, nil, err)
	} else {
		result, errs := usecase.NewFetchUserDetailUsecaseImpl(*userID, repositoryImpl.NewUserRepositoryImpl()).Run()
		helpers.Response(c, result, errs)
	}
}

func (ctrl UsersController) Create(c *gin.Context) {
	var requestJSON CreateUserRequest
	if err := c.ShouldBindJSON(&requestJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, errs := usecase.NewCreateUserUsecaseImpl(requestJSON.Name, repositoryImpl.NewUserRepositoryImpl()).Run()
	helpers.Response(c, result, errs)
}

func (ctrl UsersController) Update(c *gin.Context) {
	userID, err := user.NewUserId(helpers.ConvertToInt(c.Param("id")))
	if err.HasErrors() {
		helpers.Response(c, nil, err)
	}
	var requestJSON CreateUserRequest
	if err := c.ShouldBindJSON(&requestJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, errs := usecase.NewUpdateUserUsecaseImpl(*userID, requestJSON.Name, repositoryImpl.NewUserRepositoryImpl()).Run()
	helpers.Response(c, result, errs)
}

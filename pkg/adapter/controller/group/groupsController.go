package controller

import (
	"net/http"
	"yu-croco/ddd_on_golang/pkg/adapter/controller/helpers"
	"yu-croco/ddd_on_golang/pkg/infrastructure/repositoryImpl"
	usecase "yu-croco/ddd_on_golang/pkg/usecase/group"

	"github.com/gin-gonic/gin"
)

type GroupsController struct{}

type CreateGroupRequest struct {
	Name string `json:"name"`
}

func (ctrl GroupsController) Create(c *gin.Context) {
	var requestJSON CreateGroupRequest
	if err := c.ShouldBindJSON(&requestJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, errs := usecase.NewCreateGroupUseCaseImpl(requestJSON.Name, repositoryImpl.NewGroupRepositoryImpl()).Run()
	helpers.Response(c, result, errs)
}

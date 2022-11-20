package controller

import (
	"fmt"
	elasticsearch "yu-croco/ddd_on_golang/pkg/infrastructure/elasticSearch"

	"github.com/gin-gonic/gin"
)

type TasksController struct{}

func (ctrl TasksController) CreateIndex(c *gin.Context) {
	esAdapter := elasticsearch.NewElasticSearchAdapter()
	esAdapter.CreateIndex()
}

func (ctrl TasksController) Search(c *gin.Context) {
	esAdapter := elasticsearch.NewElasticSearchAdapter()
	fmt.Printf("esAdapter: %v\n", esAdapter)
}

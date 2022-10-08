package hunter

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UsersController struct{}


func (ctrl UsersController) Index(c *gin.Context) {
	fmt.Printf("\"こんにちンは\": %v\n", "こんにちンは")
	
}
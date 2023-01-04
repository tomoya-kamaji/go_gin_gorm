package main

import (
	"fmt"
	"os"
	cli "yu-croco/ddd_on_golang/pkg/adapter/controller/cli"
	controller "yu-croco/ddd_on_golang/pkg/adapter/controller/group"
	hunter2 "yu-croco/ddd_on_golang/pkg/adapter/controller/hunter"
	monster2 "yu-croco/ddd_on_golang/pkg/adapter/controller/monster"
	task "yu-croco/ddd_on_golang/pkg/adapter/controller/task"
	user2 "yu-croco/ddd_on_golang/pkg/adapter/controller/user"
	infrastructure2 "yu-croco/ddd_on_golang/pkg/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	db := infrastructure2.Init()
	defer db.Close()

	r := gin.Default()

	hunters := r.Group("/hunters")
	{
		huntersCtrl := hunter2.HuntersController{}
		hunters.GET("/:id", huntersCtrl.Show)
		hunters.GET("/", huntersCtrl.Index)

		hunterAttackCtrl := hunter2.HunterAttackController{}
		hunters.PUT("/:id/attack", hunterAttackCtrl.Update)

		hunterGetMaterialCtrl := hunter2.HunterGetMatrialController{}
		hunters.POST("/:id/get_material_from_monster", hunterGetMaterialCtrl.Update)
	}

	monsters := r.Group("/monsters")
	{
		monsterCtrl := monster2.Controller{}
		monsters.GET("/:id", monsterCtrl.Show)
		monsters.GET("/", monsterCtrl.Index)

		monsterAttackCtrl := monster2.MonsterAttackController{}
		monsters.PUT("/:id/attack", monsterAttackCtrl.Update)
	}

	users := r.Group("/users")
	{
		userCtrl := user2.UsersController{}
		users.GET("/", userCtrl.Index)
		users.POST("/", userCtrl.Create)
		users.PUT("/:id", userCtrl.Update)
	}

	groups := r.Group("/groups")
	{
		groupCtrl := controller.GroupsController{}
		groups.POST("/", groupCtrl.Create)
	}

	tasks := r.Group("/tasks")
	{
		taskCtrl := task.TasksController{}
		tasks.GET("/", taskCtrl.Search)
		tasks.POST("/create_index", taskCtrl.CreateIndex)
	}

	clis := r.Group("/clis")
	{
		cliCtrl := cli.ClisController{}
		clis.POST("/", cliCtrl.Run)
	}

	if err := r.Run(); err != nil {
		fmt.Printf("error occured %v", err)
		os.Exit(1)
	}
}

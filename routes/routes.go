package routes

import (
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/Harsha-S2604/to-do-list_server/service/todoservice"
)

func SetupRouter(todoDB *sql.DB) *gin.Engine{

	router := gin.Default()
	config := cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders:     []string{"Origin", "content-type"},
	}

	router.Use(cors.New(config))

	todoAPIRouter := router.Group("api/v1/todo/task")
	{
		todoAPIRouter.POST("/add", todoservice.AddTaskHandler(todoDB))
	}

	return router

}
package routes

import (
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/Harsha-S2604/to-do-list_server/service/todoservice"
	"github.com/Harsha-S2604/to-do-list_server/service/userservice"

)

func SetupRouter(todoDB *sql.DB) *gin.Engine{
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// config := cors.Config{
	// 	AllowOrigins: []string{"http://localhost:8081"},
	// 	AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin"},
	// }

	router.Use(cors.Default())

	todoAPIRouter := router.Group("api/v1/todo/task")
	{
		todoAPIRouter.GET("/tasks/:id", todoservice.GetTasksHandler(todoDB))

		todoAPIRouter.PUT("/update/:id", todoservice.UpdateTaskHandler(todoDB))

		todoAPIRouter.POST("/add", todoservice.AddTaskHandler(todoDB))

		todoAPIRouter.DELETE("/remove/:id", todoservice.RemoveTaskHandler(todoDB))
	}

	userAPIRouter := router.Group("api/v1/todo/user")
	{
		userAPIRouter.POST("/add", userservice.AddUserHandler(todoDB))
	}

	return router

}
package main

import (
	"os"
	"github.com/Harsha-S2604/to-do-list_server/config/db"
	"github.com/Harsha-S2604/to-do-list_server/routes"
)

func main() {

	todoDB, todoDBErr := db.ConnectDB()
	// redisClient := db.InitRedis()
	if todoDBErr != nil {
		panic("Database connection failed: " + todoDBErr.Error())
	} else {
		r := routes.SetupRouter(todoDB)
		port := os.Getenv("PORT") 
		r.Run(":"+port)
		defer todoDB.Close()
	}
	
	
}


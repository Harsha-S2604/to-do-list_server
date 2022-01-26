package main

import (
	"github.com/Harsha-S2604/to-do-list_server/config/db"
	"github.com/Harsha-S2604/to-do-list_server/routes"
)

func main() {

	todoDB, todoDBErr := db.ConnectDB()
	if todoDBErr != nil {
		panic("Database connection failed: " + todoDBErr.Error())
	} else {
		r := routes.SetupRouter(todoDB)
		r.Run(":8080")
		defer todoDB.Close()
	}
	
	
}


package main

import (
	"fmt"
	"github.com/Harsha-S2604/to-do-list_server/config/db"
)

func main() {

	todoDB, todoDBErr := db.ConnectDB()
	if todoDBErr != nil {
		panic(todoDBErr.Error())
	}
	
	fmt.Println(toDoDB)
}


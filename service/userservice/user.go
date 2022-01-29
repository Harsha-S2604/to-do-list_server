package userservice

import (
	"database/sql"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/Harsha-S2604/to-do-list_server/models"
	"github.com/Harsha-S2604/to-do-list_server/utility"
)

func AddUserHandler(todoDB *sql.DB) gin.HandlerFunc {

	addUser := func(ctx *gin.Context) {
		var user models.User
		ctx.ShouldBindJSON(&user)

		isValid, isValidMsg := utility.ValidateEmail(user.Email)
		if !isValid {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": isValidMsg,
			})
			return
		}

		isUserExist, isUserExistMsg := utility.CheckUserExists(user.Email, todoDB)
		if isUserExistMsg == "user exist" {
			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": isUserExistMsg,
			})
			return
		}
		if !isUserExist {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": isUserExistMsg,
			})
			return
		} 

		addUserQuery := `INSERT INTO users VALUES($1);`
		_, addUserQueryErr := todoDB.Exec(addUserQuery, user.Email)
		if addUserQueryErr != nil {
			log.Println(addUserQueryErr.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Sorry something went wrong. Please try again later",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User registered successfully",
		})
		
	}

	return gin.HandlerFunc(addUser)
}
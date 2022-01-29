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

		var id int
		tx, beginErr := todoDB.Begin()
		if beginErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"success": false,
				"message": "Sorry, Something went wrong. Please refresh the page or try again later.",
			})
			return
		}

		{
			addUserQueryStmt := `INSERT INTO users(email) VALUES($1) RETURNING user_id;`
			addUserQuery, addUserQueryErr := tx.Prepare(addUserQueryStmt)
			if addUserQueryErr != nil {
				log.Println(addUserQueryErr.Error())
				ctx.JSON(http.StatusOK, gin.H{
					"success": false,
					"message": "Sorry something went wrong. Please try again later",
				})
				return
			}
		

			defer addUserQuery.Close()

			addUserQueryErr = addUserQuery.QueryRow(
				user.Email,
			).Scan(&id)

			if addUserQueryErr != nil {
				log.Println(addUserQueryErr.Error())
				ctx.JSON(http.StatusOK, gin.H{
					"success": false,
					"message": "Sorry something went wrong. Please try again later",
				})
				return
			}
		}

		{

			commitErr := tx.Commit()
			if commitErr != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"success": false,
					"message": "Sorry, Something went wrong. Please refresh the page or try again later.",
				})
				return
			}

		}

		

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User registered successfully",
			"id": id,
		})
		
	}

	return gin.HandlerFunc(addUser)
}
package todoservice

import (
	"fmt"
	"net/http"
	"database/sql"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Harsha-S2604/to-do-list_server/models"

)

func AddTaskHandler(todoDB *sql.DB) gin.HandlerFunc {
	
	addTask := func(ctx *gin.Context) {
		var task models.Task
		ctx.ShouldBindJSON(&task)
		isDataValid, message := validateData(task)

		if !isDataValid {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": message,
			})
			return
		}

		isUserExists, isUserExistsMessage := checkUserExists(task.User.Email, todoDB)
		if !isUserExists {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": isUserExistsMessage,
			})
			return
		}

		addTaskQuery := `INSERT INTO todo_list (task_name, is_completed, email) VALUES($1, $2, $3)`
		_, addTaskQueryErr := todoDB.Exec(addTaskQuery, task.TaskName, task.IsCompleted, task.User.Email)
		if addTaskQueryErr != nil {
			fmt.Println(addTaskQueryErr.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Sorry something went wrong. Please try again later",
			})
			return
		}
 
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Task added successfully.",
		})
	}
	
	return gin.HandlerFunc(addTask)
}

func RemoveTaskHandler(todoDB *sql.DB) gin.HandlerFunc {

	removeTask := func(ctx *gin.Context) {
		taskId, taskIdErr := strconv.Atoi(ctx.Params.ByName("id"))

		if taskIdErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Sorry something went wrong. Please try again later.",
			})
			return
		}

		if taskId == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Task does not exist",
			})
			return
		}

		isTaskExist, isTaskExistMsg := checkTaskExist(taskId, todoDB)
		if !isTaskExist {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": isTaskExistMsg,
			})
			return
		}

		removeTaskQuery := `DELETE FROM todo_list WHERE task_id=$1`
		_, removeTaskQueryErr := todoDB.Exec(removeTaskQuery, taskId)
		if removeTaskQueryErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Sorry something went wrong. Please try again later.",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Task removed successfully",
		})
	}

	return gin.HandlerFunc(removeTask)
}

func validateData(task models.Task) (bool, string) {
	if task.User.Email == "" {
		return false, "email is required"
	}

	if task.TaskName == "" {
		return false, "task name is required"
	}

	var re = regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$`)
	if !re.MatchString(task.User.Email) {
		return false, "not a valid email"
	} 


	return true, ""
}

func checkTaskExist(taskId int, todoDB *sql.DB) (bool, string) {
	var taskFromDB string
	taskExistsQuery := `SELECT task_name from todo_list WHERE task_id=$1;`
	row := todoDB.QueryRow(taskExistsQuery, taskId)
	err := row.Scan(&taskFromDB)
	switch err {
    case sql.ErrNoRows:
        return false, "task does not exist"
    case nil:
        return true, ""
    default:
        return false, "Sorry something went wrong. Please try again later."
    }
}

func checkUserExists(email string, todoDB *sql.DB) (bool, string) {
	var emailFromDB string
	userExistsQuery := `SELECT email from users WHERE email=$1;`
	row := todoDB.QueryRow(userExistsQuery, email)
	err := row.Scan(&emailFromDB)
	switch err {
    case sql.ErrNoRows:
        return false, "user does not exist"
    case nil:
        return true, ""
    default:
        return false, "Sorry something went wrong. Please try again later."
    }
}
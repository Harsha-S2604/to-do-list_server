package todoservice

import (
	"fmt"
	"context"
	"net/http"
	"database/sql"
	"regexp"
	// "time"
	"strconv"
	// "encoding/json"

	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v8"
	"github.com/Harsha-S2604/to-do-list_server/models"

)

var contextGlobal = context.Background()

func GetTasksHandler(todoDB *sql.DB) gin.HandlerFunc {

	getTasks := func(ctx *gin.Context) {
		var tasks []models.Task
		userIdStr := ctx.Params.ByName("id")
		
		userId, userIdErr := strconv.Atoi(userIdStr)
		if userIdErr != nil {
			fmt.Println("useriderr", userIdErr.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Sorry something went wrong. Please try again later.",
			})
			return
		}

		isUserExists, isUserExistsMessage := checkUserExists(userId, todoDB)
		if !isUserExists {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"data": nil,
				"message": isUserExistsMessage,
			})
			return	
		}

		queryParams := ctx.Request.URL.Query()
		limitStr, ok := queryParams["limit"]
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"data": nil,
				"message": "invalid request",
			})
			return
		}
		limit, limitErr := strconv.Atoi(limitStr[0])
		if limitErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"data": nil,
				"message": "Sorry something went wrong. Please try again later.",
			})
			return
		}

		offsetStr, ok := queryParams["offset"]
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"data": nil,
				"message": "invalid request",
			})
			return
		}

		offset, offsetErr := strconv.Atoi(offsetStr[0])
		if offsetErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"data": nil,
				"message": "Sorry something went wrong. Please try again later.",
			})
			return
		}

		getQuery := `SELECT * FROM todo_list WHERE user_id=$1 ORDER BY task_id DESC LIMIT $2 OFFSET $3`
		rows, rowsErr := todoDB.Query(getQuery, userId, limit, offset - 1)
		if rowsErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"data": nil,
				"message": "Sorry something went wrong. Please try again later.",
			})
			return
		}
		for rows.Next() {
			var task models.Task
	
			rowScanErr := rows.Scan(&task.TaskId, &task.TaskName, &task.IsCompleted, &task.User.UserId)
	
			if rowScanErr != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"success": false,
					"data": nil,
					"message": "Sorry something went wrong. Please try again later.",
				})
				return
			}
	
			tasks = append(tasks, task)
	
		}


		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data": tasks,
		})
		
	}

	return gin.HandlerFunc(getTasks)
}

func AddTaskHandler(todoDB *sql.DB) gin.HandlerFunc {
	
	addTask := func(ctx *gin.Context) {
		var task models.Task
		ctx.ShouldBindJSON(&task)
		isDataValid, message := validateData(task)
		fmt.Println(task)
		if !isDataValid {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": message,
			})
			return
		}

		isUserExists, isUserExistsMessage := checkUserExists(task.User.UserId, todoDB)
		if !isUserExists {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": isUserExistsMessage,
			})
			return
		}

		addTaskQuery := `INSERT INTO todo_list (task_name, is_completed, user_id) VALUES($1, $2, $3)`
		_, addTaskQueryErr := todoDB.Exec(addTaskQuery, task.TaskName, task.IsCompleted, task.User.UserId)
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

		// queryParams := ctx.Request.URL.Query()
		// offsetStr, ok := queryParams["offset"]
		// if !ok {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"success": false,
		// 		"data": nil,
		// 		"message": "invalid request",
		// 	})
		// 	return
		// }

		// offset, offsetErr := strconv.Atoi(offsetStr[0])
		// if offsetErr != nil {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"success": false,
		// 		"data": nil,
		// 		"message": "Sorry something went wrong. Please try again later.",
		// 	})
		// 	return
		// }

		// userIdStr, ok := queryParams["userId"]
		// if !ok {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"success": false,
		// 		"data": nil,
		// 		"message": "invalid request",
		// 	})
		// 	return
		// }
		
		// userId, userIdErr := strconv.Atoi(userIdStr[0])
		// if userIdErr != nil {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"success": false,
		// 		"data": nil,
		// 		"message": "Sorry something went wrong. Please try again later.",
		// 	})
		// 	return
		// }

		// redisCachedName := "todo_list_"+userIdStr[0]+"_"+offsetStr[0]

		removeTaskQuery := `DELETE FROM todo_list WHERE task_id=$1`
		_, removeTaskQueryErr := todoDB.Exec(removeTaskQuery, taskId)
		if removeTaskQueryErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Sorry something went wrong. Please try again later.",
			})
			return
		}
		// toBeCachedTodoList, toBeCachedTodoListErr := getTodoLists(todoDB, userId, offset)
		// marshaledTodoList, _  := json.Marshal(toBeCachedTodoList)
		// if toBeCachedTodoListErr == nil {
		// 	_ = redisClient.Set(ctx, redisCachedName, marshaledTodoList, 100 * time.Second).Err()
		// }

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Task removed successfully",
		})
	}

	return gin.HandlerFunc(removeTask)
}

func UpdateTaskHandler(todoDB *sql.DB) gin.HandlerFunc {
	
	updateTask := func(ctx *gin.Context) {
		
		taskId, taskIdErr := strconv.Atoi(ctx.Params.ByName("id"))

		if taskIdErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Sorry something went wrong. Please try again later.",
			})
			return
		}
		form, formErr := ctx.MultipartForm()
		if formErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Sorry something went wrong. Please try again later.",
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
		isCompletedArr := form.Value["isCompleted"]
		var isCompleted interface{}
		isCompleted = isCompletedArr[0]

		// queryParams := ctx.Request.URL.Query()
		// offsetStr, ok := queryParams["offset"]
		// if !ok {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"success": false,
		// 		"data": nil,
		// 		"message": "invalid request",
		// 	})
		// 	return
		// }

		// offset, offsetErr := strconv.Atoi(offsetStr[0])
		// if offsetErr != nil {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"success": false,
		// 		"data": nil,
		// 		"message": "Sorry something went wrong. Please try again later.",
		// 	})
		// 	return
		// }

		// userIdStr, ok := queryParams["userId"]
		// if !ok {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"success": false,
		// 		"data": nil,
		// 		"message": "invalid request",
		// 	})
		// 	return
		// }
		
		// userId, userIdErr := strconv.Atoi(userIdStr[0])
		// if userIdErr != nil {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"success": false,
		// 		"data": nil,
		// 		"message": "Sorry something went wrong. Please try again later.",
		// 	})
		// 	return
		// }

		// redisCachedName := "todo_list_"+userIdStr[0]+"_"+offsetStr[0]

		updateQuery := `UPDATE todo_list SET is_completed=$1 WHERE task_id=$2;`
		_, updateErr := todoDB.Exec(updateQuery, isCompleted, taskId)
		if updateErr != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Sorry something went wrong. Please try again later.",
			})
			return
		}

		// toBeCachedTodoList, toBeCachedTodoListErr := getTodoLists(todoDB, userId, offset)
		// marshaledTodoList, _  := json.Marshal(toBeCachedTodoList)
		// if toBeCachedTodoListErr == nil {
		// 	_ = redisClient.Set(ctx, redisCachedName, marshaledTodoList, 100 * time.Second).Err()
		// }


		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "task updated successfully",
		})
	}

	return gin.HandlerFunc(updateTask)
}

func validateData(task models.Task) (bool, string) {
	if task.User.UserId == 0 {
		return false, "not a valid request. please try again later."
	}

	if task.TaskName == "" {
		return false, "task name is required"
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

func checkUserExists(userId int, todoDB *sql.DB) (bool, string) {
	var emailFromDB string
	userExistsQuery := `SELECT email from users WHERE user_id=$1;`
	row := todoDB.QueryRow(userExistsQuery, userId)
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

func validateEmail(email string) (bool, string) {
	fmt.Println(email)
	var re = regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$`)
	if !re.MatchString(email) {
		return false, "not a valid email"
	}

	return true, ""
}

func getTodoLists(todoDB *sql.DB, userId, offset int) ([]models.Task, error){
	getQuery := `SELECT * FROM todo_list WHERE user_id=$1 ORDER BY task_id DESC LIMIT 5 OFFSET $2`
	rows, _ := todoDB.Query(getQuery, userId, offset - 1)
	var tasks []models.Task
	for rows.Next() {
		var task models.Task

		rowScanErr := rows.Scan(&task.TaskId, &task.TaskName, &task.IsCompleted, &task.User.UserId)

		if rowScanErr != nil {
			return nil, rowScanErr
		}

		tasks = append(tasks, task)

	}
	return tasks, nil

}
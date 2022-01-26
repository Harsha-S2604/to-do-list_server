package todoservice

import (
	"testing"
	"database/sql"
	"net/http"
  	"net/http/httptest"
	"os"
	"log"
    "bytes"
	"encoding/json"

	"github.com/Harsha-S2604/to-do-list_server/config/db"
	// "github.com/Harsha-S2604/to-do-list_server/models"
	"github.com/Harsha-S2604/to-do-list_server/service/todoservice"
	"github.com/gin-gonic/gin"
)

var todoDB *sql.DB
var todoDBErr error

func TestMain(m *testing.M) {
	todoDB, todoDBErr = db.ConnectDB()
	if todoDBErr != nil {
		panic(todoDBErr)
	}
	exitVal := m.Run()	

    os.Exit(exitVal)
}

func TestGetTasks(t *testing.T) {

	// var resp []models.Task
	gin.SetMode(gin.TestMode)

	r := gin.Default()
    r.GET("/", todoservice.GetTasksHandler(todoDB))
	body := bytes.NewBuffer([]byte("{\"email\":\"arix2604@gmail.com\"}"))

	req, reqErr := http.NewRequest(http.MethodGet, "/?limit=5&offset=1", body)
	req.Header.Set("Content-Type", "application/json")
    if reqErr != nil {
        t.Fatalf("Couldn't create request: %v\n", reqErr)
    }

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
    if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	respBody := map[string]interface{}{}
    if err := json.Unmarshal([]byte(w.Body.String()), &respBody); err != nil {
        panic("unmarshaling response body returned error")
    }

	expected := true

	if respBody["success"] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
		respBody["success"], expected)
	}
	

}

func TestAddTasks(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
    r.POST("/", todoservice.AddTaskHandler(todoDB))
	body := bytes.NewBuffer([]byte("{\"user\":{\"email\":\"arix2604@gmail.com\"},\"taskName\":\"task_2\",\"isCompleted\":\"false\"}"))

	req, reqErr := http.NewRequest(http.MethodPost, "/", body)
	req.Header.Set("Content-Type", "application/json")
    if reqErr != nil {
        t.Fatalf("Couldn't create request: %v\n", reqErr)
    }

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
    if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	respBody := map[string]interface{}{}
    if err := json.Unmarshal([]byte(w.Body.String()), &respBody); err != nil {
        panic("unmarshaling response body returned error")
    }

	expected := true
	log.Println(respBody)

	if respBody["success"] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
		respBody["success"], expected)
	}
	

}

func TestRemoveTasks(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
    r.DELETE("/:id", todoservice.RemoveTaskHandler(todoDB))

	req, reqErr := http.NewRequest(http.MethodDelete, "/5", nil)
	req.Header.Set("Content-Type", "application/json")
    if reqErr != nil {
        t.Fatalf("Couldn't create request: %v\n", reqErr)
    }

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
    if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	respBody := map[string]interface{}{}
    if err := json.Unmarshal([]byte(w.Body.String()), &respBody); err != nil {
        panic("unmarshaling response body returned error")
    }

	expected := true
	log.Println(respBody)

	if respBody["success"] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
		respBody["success"], expected)
	}
	

}
package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Fumiya-Tahara/todo-app/generated/openapi"
	"github.com/Fumiya-Tahara/todo-app/internal/domain"
	"github.com/Fumiya-Tahara/todo-app/internal/infrastracture"
)

type handler struct{}

func NewHandler() openapi.ServerInterface {
	return &handler{}
}

func (h handler) GetTaskList(w http.ResponseWriter, r *http.Request) {
	client := infrastracture.NewStorage()
	defer client.DB.Close()
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		taskHandler := &domain.TaskStorage{
			DB: client.DB,
		}
		data, err := taskHandler.GetTasks()
		if err != nil {
			log.Fatal(err)
		}
		w.Write(data)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func (h handler) GetSpecificTask(w http.ResponseWriter, r *http.Request, id int) {
	client := infrastracture.NewStorage()
	defer client.DB.Close()
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		taskHandler := &domain.TaskStorage{
			DB: client.DB,
		}
		data, err := taskHandler.GetSpecificTask(id)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(data)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func (h handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	client := infrastracture.NewStorage()
	defer client.DB.Close()
	// if r.Method == "POST" {
	// 	w.WriteHeader(http.StatusOK)

	// }
	title := "testTitle"
	content := "testContent"
	deadline := time.Now()

	taskHandler := &domain.TaskStorage{
		DB: client.DB,
	}

	taskHandler.CreateTasks(title, content, deadline)
}

func (h handler) UpdateTask(w http.ResponseWriter, r *http.Request, id int) {
	client := infrastracture.NewStorage()
	defer client.DB.Close()

	title := "更新したタイトル"
	content := "更新した内容"
	isCompleted := false
	deadline := time.Now()

	taskHandler := &domain.TaskStorage{
		DB: client.DB,
	}

	taskHandler.UpdateTask(id, title, content, isCompleted, deadline)

}

func (h handler) DeleteTask(w http.ResponseWriter, r *http.Request, id int) {
	client := infrastracture.NewStorage()
	defer client.DB.Close()

	taskHandler := &domain.TaskStorage{
		DB: client.DB,
	}

	taskHandler.DeleteTask(id)
}

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	h := NewHandler()
	trimPath := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if trimPath == "" {
		h.GetTaskList(w, r)
		return
	}

	id, err := strconv.Atoi(trimPath)
	if err != nil {
		log.Fatal(err)
	}
	h.GetSpecificTask(w, r, id)
}

func HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	h := NewHandler()
	trimPath := strings.TrimPrefix(r.URL.Path, "/update-task/")
	id, err := strconv.Atoi(trimPath)
	if err != nil {
		log.Fatal(err)
	}
	h.UpdateTask(w, r, id)
}

func HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	h := NewHandler()
	trimPath := strings.TrimPrefix(r.URL.Path, "/delete-task/")
	id, err := strconv.Atoi(trimPath)
	if err != nil {
		log.Fatal(err)
	}
	h.DeleteTask(w, r, id)
}

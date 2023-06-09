package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

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

func (h handler) GetTasksId(w http.ResponseWriter, r *http.Request, id int) {
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

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	h := NewHandler()
	trimPath := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if trimPath == "" {
		h.GetTaskList(w, r)
		return
	}

	id, err := strconv.Atoi(trimPath)
	if err != nil {
		log.Println(err)
		return
	}
	h.GetTasksId(w, r, id)
}

func (h handler) PostCreateTask(w http.ResponseWriter, r *http.Request) {
	return
}
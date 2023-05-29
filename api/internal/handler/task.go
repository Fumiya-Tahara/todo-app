package handler

import (
	"log"
	"net/http"
)

type handler struct{}

func NewHandler() ServerInterface {
	return &handler
}

func (h handler) GetTaskList(w http.ResponseWriter, r *http.Request) {
	log.Println("成功")
	return nil
}

func (h handler) GetTasksId(w http.ResponseWriter, r *http.Request, id int) {
	return nil
}

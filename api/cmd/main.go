package main

import (
	"fmt"
	"net/http"

	"github.com/Fumiya-Tahara/todo-app/internal/handler"
)

func main() {
	h := handler.NewHandler()
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) { h.GetTaskList(w, r) })

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

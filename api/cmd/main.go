package main

import (
	"fmt"
	"net/http"

	"github.com/Fumiya-Tahara/todo-app/handler"
)

func main() {
	http.HandleFunc("/tasks", handler.GetTaskList())

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

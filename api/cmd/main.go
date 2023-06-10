package main

import (
	"fmt"
	"net/http"

	"github.com/Fumiya-Tahara/todo-app/internal/handler"
)

func main() {
	http.HandleFunc("/tasks/", handler.HandleTasks)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

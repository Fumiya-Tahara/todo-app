package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fumiya-Tahara/todo-app/internal/domain"
	"github.com/Fumiya-Tahara/todo-app/internal/infrastracture"
)

func main() {
	client := infrastracture.NewStorage()
	defer client.DB.Close()
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		taskHandler := &domain.TaskStorage{
			DB: client.DB,
		}
		data, err := taskHandler.GetTasks()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

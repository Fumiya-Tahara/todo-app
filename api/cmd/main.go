package main

import (
	"fmt"
	"net/http"

	"github.com/Fumiya-Tahara/todo-app/internal/infrastracture"
)

func main() {
	client := infrastracture.NewStorage()
	defer client.DB.Close()
	// http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// 	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// 	if r.Method == "GET" {
	// 		w.WriteHeader(http.StatusOK)
	// 		taskHandler := &domain.TaskStorage{
	// 			DB: client.DB,
	// 		}
	// 		data, err := taskHandler.GetTasks()
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Println(string(data))
	// 		w.Write(data)
	// 	} else {
	// 		w.WriteHeader(http.StatusMethodNotAllowed)
	// 		w.Write([]byte("Method not allowed"))
	// 	}
	// })

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

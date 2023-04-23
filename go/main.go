package main

import (
	"fmt"

	"github.com/Fumiya-Tahara/todo-app/infrastracture"
)

func main() {
	fmt.Printf("hello world")
	infrastracture.ConnectSql()
}

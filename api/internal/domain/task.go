package domain

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	IsCompleted bool      `json:"isCompleted"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Tasks []Task

type TaskStorage struct {
	DB *sql.DB
}

func (storage *TaskStorage) GetTasks() ([]byte, error) {
	query := "SELECT id, title, content, is_completed, deadline, created_at, updated_at FROM tasks"
	rows, err := storage.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.IsCompleted, &task.Deadline, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}

	return jsonTasks, nil
}

func (storage *TaskStorage) GetSpecificTask(id int) ([]byte, error) {
	strId := strconv.Itoa(id)
	query := "SELECT id, title, content, is_completed, deadline FROM tasks WHERE id = " + strId
	rows, err := storage.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.IsCompleted, &task.Deadline)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	log.Println(tasks)
	jsonTask, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}

	return jsonTask, nil
}

func (storage *TaskStorage) CreateTasks(title string, content string, deadline time.Time) {
	query := "INSERT INTO tasks (title, content, deadline) VALUES (?, ?, ?)"
	result, err := storage.DB.Exec(query, title, content, deadline)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}
}

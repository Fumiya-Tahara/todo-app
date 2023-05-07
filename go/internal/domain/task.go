package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID          int
	Title       string
	Content     string
	IsCompleted bool
	Deadline    *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tasks []Task

type TaskStorage struct {
	DB *sql.DB
}

func (storage *TaskStorage) GetTasks() ([]Task, error) {
	rows, err := storage.DB.Query("SELECT id, title, content, is_completed, deadline, created_at, updated_at FROM tasks")
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
	return tasks, nil
}

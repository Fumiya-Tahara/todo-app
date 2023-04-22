package model

import (
	"time"
)

type Task struct {
	ID          int
	UserId      int
	Title       string
	Content     string
	IsCompleted bool
	Deadline    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tasks []Task

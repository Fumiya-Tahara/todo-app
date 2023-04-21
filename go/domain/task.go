package domain

import "time"

type Task struct {
	ID          int
	UserId      int
	Title       string
	Content     string
	IsCompleted bool
	Deadline    time.Time
}

type Tasks []Task

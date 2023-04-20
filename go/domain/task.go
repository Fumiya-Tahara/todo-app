package domain

type Task struct {
	ID      int
	UserId  int
	Title   string
	Content string
	Status  int
}

type Tasks []Task

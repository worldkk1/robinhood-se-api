package domain

type TaskStatus string

const (
	ToDo TaskStatus = "to_do"
	InProgress TaskStatus = "in_progress"
	Done TaskStatus = "done"
)

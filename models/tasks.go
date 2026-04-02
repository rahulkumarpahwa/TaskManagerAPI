package models

import "time"

type Status string

const (
	Completed    Status = "Completed"
	Pending      Status = "Pending"
	UnderProcess Status = "UnderProcess"
	Skipped      Status = "Skipped"
	UnCompleted      Status = "UnCompleted"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

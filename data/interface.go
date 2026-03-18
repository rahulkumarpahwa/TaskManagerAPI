package data

import (
	"TaskManager/models"
)

type TaskRepositaryModel interface {
	CreateTask(task models.CreateTask) int
	GetTasks() ([]models.Task, error)
	UpdateTask(id int, title string, description string, status models.Status) (models.Task, error)
	DeleteTask(id int) (int, error)
}

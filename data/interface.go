package data

import (
	"TaskManager/models"
)

type TaskRepositaryModel interface {
	CreateTask(models.CreateTask)
	GetTasks() []models.Task
}

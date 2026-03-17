package data

import (
	"TaskManager/models"
	"database/sql"
	"time"
)

type TaskRepositary struct {
	DB *sql.DB
}

func (tr *TaskRepositary) CreateTask(task models.Task) {

}

func (tr *TaskRepositary) GetTasks() []models.Task {

	var tasks []models.Task

	tasks = []models.Task{{
		ID:          34,
		Title:       "Hello",
		Description: "FFFF",
		Status:      "Completed",
		CreatedAt:   time.Now(),
	}}

	return tasks
}

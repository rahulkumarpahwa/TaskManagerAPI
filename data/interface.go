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

type UserRepositaryModel interface {
	Register(user models.CreateUser) (bool, error)
	FindUserById(id int) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	DeleteUser(id int) error
	FavoriteTasks(task_id int, user_id int)
}

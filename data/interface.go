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
	Register(string, string, string) (int, bool, error)
	Authenticate(string, string) (int, bool, error)
	// FindUserById(int) (models.User, error)
	// FindUserByEmail(string) (models.User, error)
	// DeleteUser(int) error
	// FavoriteTasks(int, int)
}

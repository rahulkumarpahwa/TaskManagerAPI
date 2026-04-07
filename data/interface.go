package data

import (
	"TaskManager/models"
)

type TaskRepositaryModel interface {
	CreateTask(models.CreateTask) int
	GetTasks() ([]models.Task, error)
	UpdateTask(int, string, string, models.Status, int, bool) (models.Task, bool, error)
	DeleteTask(int, int) (models.Task, bool, error)
}

type UserRepositaryModel interface {
	Register(string, string, string) (int, bool, error)
	Authenticate(string, string) (models.User, bool, error)
	// FindUserById(int) (models.User, error)
	// FindUserByEmail(string) (models.User, error)
	// DeleteUser(int) error
	// FavoriteTasks(int, int)
}

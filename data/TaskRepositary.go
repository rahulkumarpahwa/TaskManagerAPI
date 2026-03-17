package data

import (
	"TaskManager/models"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

type TaskRepositary struct {
	DB *sql.DB
}

const defaultLimit = 20

func (tr *TaskRepositary) CreateTask(task models.CreateTask) int {
	query := `INSERT INTO TASKS (title, description, status) VALUES ($1, $2, $3)`

	result, err := tr.DB.Exec(query, task.Title, task.Description, task.Status)

	if err != nil {
		log.Fatalf("Task can't be created : %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("could not get affected rows: %v", err)
	}
	return int(rowsAffected)
}

func (tr *TaskRepositary) GetTasks() []models.Task {
	query := `SELECT * FROM TASKS`
	rows, err := tr.DB.Query(query, defaultLimit)
	if err != nil {
		log.Fatalf("could not get tasks: %v", err)
	}
	defer rows.Close()
	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Description, &task.Status, &task.CreatedAt); err != nil {
			log.Fatalf("Can't retrive the tasks row : %v", err)
			return []models.Task{}
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func (tr *TaskRepositary) UpdateTask(id int, title string, description string, status models.Status) models.Task {
	query := `UPDATE Tasks SET title = $1, description = $2, status = $3 WHERE id = $4`
	result, err := tr.DB.Exec(query, title, description, status)
	if err != nil {
		log.Fatalf("Task can't be created : %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("could not get affected rows: %v", err)
	}
	if int(rowsAffected) > 0 {
		query := `SELECT * FROM Tasks WHERE id = $1`
		row := tr.DB.QueryRow(query)
		var task models.Task
		if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt); err != nil {
			log.Fatalf("Can't get the updated task, %v", err)
		}
		return task
	} else {
		return models.Task{}
	}
}

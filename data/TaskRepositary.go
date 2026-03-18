package data

import (
	"TaskManager/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
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

func (tr *TaskRepositary) GetTasks() ([]models.Task, error) {
	query := `SELECT id, title, description, status, created_at FROM tasks ORDER BY id LIMIT $1`
	rows, err := tr.DB.Query(query, defaultLimit)
	if err != nil {
		return nil, fmt.Errorf("could not get tasks: %w", err)
	}
	defer rows.Close()
	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt); err != nil {
			return nil, fmt.Errorf("can't retrieve task row: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("task rows iteration failed: %w", err)
	}

	return tasks, nil
}

func (tr *TaskRepositary) UpdateTask(id int, title string, description string, status models.Status) (models.Task, error) {
	query := `UPDATE Tasks SET title = $1, description = $2, status = $3 WHERE id = $4`
	result, err := tr.DB.Exec(query, title, description, status, id)
	if err != nil {
		return models.Task{}, fmt.Errorf("could not get updated task: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Task{}, fmt.Errorf("could not get any rows affected: %w", err)
	}
	if int(rowsAffected) > 0 {
		query := `SELECT id, title, description, status, created_at FROM Tasks WHERE id = $1`
		row := tr.DB.QueryRow(query, id)
		var task models.Task
		if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt); err != nil {
			return models.Task{}, fmt.Errorf("could not updated task: %w", err)
		}
		return task, nil
	} else {
		return models.Task{}, nil
	}
}

func (tr *TaskRepositary) DeleteTask(id int) (int, error) {
	query := `DELETE FROM Tasks WHERE id = $1`
	result, err := tr.DB.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("could not get delete task: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not get any rows affected: %w", err)
	}
	if int(rowsAffected) > 0 {
		return int(rowsAffected), nil
	} else {
		return 0, nil
	}
}

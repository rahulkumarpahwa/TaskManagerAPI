package data

import (
	"TaskManager/models"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type TaskRepositary struct {
	DB *sql.DB
}

const defaultLimit = 20

func (tr *TaskRepositary) CreateTask(task models.CreateTask) int {
	query := `INSERT INTO TASKS (title, description, status, user_id, is_favorite) VALUES ($1, $2, $3, $4, $5)`

	result, err := tr.DB.Exec(query, task.Title, task.Description, task.Status, task.UserId, task.IsFavorite)

	if err != nil {
		log.Printf("Task can't be created : %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("could not get affected rows: %v", err)
	}
	return int(rowsAffected)
}

func (tr *TaskRepositary) GetTasks() ([]models.Task, error) {
	query := `SELECT id, title, description, status, created_at, modified_at, user_id, is_favorite FROM tasks ORDER BY id LIMIT $1`
	rows, err := tr.DB.Query(query, defaultLimit)
	if err != nil {
		return nil, fmt.Errorf("could not get tasks: %w", err)
	}
	defer rows.Close()
	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.ModifiedAt, &task.UserId, &task.IsFavorite); err != nil {
			return nil, fmt.Errorf("can't retrieve task row: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("task rows iteration failed: %w", err)
	}

	return tasks, nil
}

func (tr *TaskRepositary) UpdateTask(id int, title string, description string, status models.Status, userId int, isFavorite bool) (models.Task, bool, error) {
	query := `UPDATE Tasks SET title = $1, description = $2, status = $3, is_favorite = $4 WHERE id = $5 AND user_id = $6`
	_, err := tr.DB.Exec(query, title, description, status, isFavorite, id, userId)
	if err != nil {
		return models.Task{}, false, fmt.Errorf("could not get updated task: %w", err)
	}

	// updating the modified_at:
	updateQuery := `UPDATE Tasks SET modified_at = $1 WHERE id = $2 RETURNING id, title, description, status, created_at, modified_at, user_id, is_favorite`

	var updatedTask models.Task
	err = tr.DB.QueryRow(updateQuery, time.Now(), id).Scan(&updatedTask.ID, &updatedTask.Title, &updatedTask.Description, &updatedTask.Status, &updatedTask.CreatedAt, &updatedTask.ModifiedAt, &updatedTask.UserId, &updatedTask.IsFavorite)

	if err == sql.ErrNoRows {
		return models.Task{}, false, fmt.Errorf("Not able to update the modified time : %v", err)
	}
	if err != nil {
		return models.Task{}, false, fmt.Errorf("Error while updating the modified time : %v", err)
	}

	return updatedTask, true, nil
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

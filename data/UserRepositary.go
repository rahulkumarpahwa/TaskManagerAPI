package data

import (
	"TaskManager/models"
	"database/sql"
	"errors"
	"log"
)

type UserRepositary struct {
	DB *sql.DB
}

func (ur *UserRepositary) Register(Name string, Email string, Password string) (bool, error) {
	query := `INSERT INTO USER (username, password, email) VALUES ($1, $2, $3)`
	result, err := ur.DB.Exec(query, Name, Email, Password)
	if err != nil {
		log.Printf("User can't be inserted : %v", err)
		return false, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		log.Printf("User Rows can't be inserted : %v", err)
		return false, err
	}
	return true, nil
}

func (ur *UserRepositary) FindUserById(id int) (user models.User, error error) {
	query := `SELECT id, username, password, email, created_at FROM USER WHERE id = $1`
	row := ur.DB.QueryRow(query, id)

	var u models.User
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		log.Printf("User Row can't be Scaned : %v", err)
		return models.User{}, err
	}
	return u, nil
}

func (ur *UserRepositary) FindUserByEmail(email string) (user models.User, error error) {
	query := `SELECT id, username, password, email, created_at FROM USER WHERE email = $1`
	row := ur.DB.QueryRow(query, email)

	var u models.User
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		log.Printf("User Row can't be Scaned : %v", err)
		return models.User{}, err
	}
	return u, nil
}

func (ur *UserRepositary) DeleteUser(id int) (int, error) {
	query := `DELETE FROM USER WHERE id = $1`
	result, err := ur.DB.Exec(query, id)
	if err != nil {
		log.Printf("User can't be deleted : %v", err)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("User Rows can't be deleted : %v", err)
		return 0, err
	}
	if int(rowsAffected) > 0 {
		return int(rowsAffected), nil
	} else {
		return 0, errors.New("No User Rows deletion affected!")
	}
}

func (ur *UserRepositary) FavoriteTasks(task_id int, user_id int) (int, error) {
	query := `INSERT INTO task_user (task_id, user_id) VALUES ($1, $2)`

	result, err := ur.DB.Exec(query, task_id, user_id)
	if err != nil {
		log.Printf("favorite task can't be inserted : %v", err)
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("favorite task rows can't be affected : %v", err)
		return 0, err

	}
	return int(rowsAffected), nil
}

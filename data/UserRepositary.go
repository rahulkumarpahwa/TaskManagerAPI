package data

import (
	"TaskManager/models"
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserRepositary struct {
	DB *sql.DB
}

func (ur *UserRepositary) Register(Name string, Email string, Password string) (int, bool, error) {

	if Name == "" || Email == "" || Password == "" {
		log.Print("Invalid resgistration! Requirement Missing!")
		return -1, false, nil
	}

	// check if the user exists already
	var exists bool
	err := ur.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`, Email).Scan(&exists)
	if err != nil {
		log.Printf("Failed to check user : %v", err)
		return -1, false, err
	}
	if exists {
		log.Printf("User Already Exists with email : %s", Email)
		return -1, false, errors.New("User Already Exists!")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return -1, false, err
	}

	var id int
	query := `INSERT INTO USERS (name, password_hashed, email) VALUES ($1, $2, $3) RETURNING id`
	err = ur.DB.QueryRow(query, Name, string(hashedPassword), Email).Scan(&id)
	if err != nil {
		log.Printf("User can't be inserted : %v", err)
		return -1, false, err
	}
	return id, true, nil
}

func (ur *UserRepositary) FindUserById(id int) (user models.User, error error) {
	query := `SELECT id, name, password_hashed, email, time_created FROM users WHERE id = $1`
	row := ur.DB.QueryRow(query, id)

	var u models.User
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.CreatedAt)
	if err != nil {
		log.Printf("User Row can't be Scaned : %v", err)
		return models.User{}, err
	}
	return u, nil
}

func (ur *UserRepositary) FindUserByEmail(email string) (user models.User, error error) {
	query := `SELECT id, name, password_hashed, email, time_created FROM users WHERE email = $1`
	row := ur.DB.QueryRow(query, email)

	var u models.User
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.CreatedAt)
	if err != nil {
		log.Printf("User Row can't be Scaned : %v", err)
		return models.User{}, err
	}
	return u, nil
}

func (ur *UserRepositary) DeleteUser(id int) (int, error) {
	query := `DELETE FROM users WHERE id = $1`
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

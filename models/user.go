package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name  string    `json:"name"`
	Password  string    `json:"password_hashed"`
	Email     string    `json:"email"`
	LastLogin time.Time `json:"last_login"`
	TimeCreated time.Time `json:"time_created"`
	TimeConfirmed time.Time `json:"time_confirmed"`
	TimeDeleted time.Time `json:"time_deleted"`
}

type CreateUser struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

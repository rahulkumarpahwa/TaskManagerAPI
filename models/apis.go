package models

type UserRequest struct {
	Name string 
	Email string
	Password string
}

type UserResponse struct{
	Success bool
	Message string
	
}

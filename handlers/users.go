package handlers

import (
	"TaskManager/data"
	"TaskManager/models"
	"TaskManager/token"
	"encoding/json"
	"net/http"
	"time"
)

type UserHandlers struct {
	Storage data.UserRepositaryModel
}

func (h *UserHandlers) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqBody models.UserRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Can't Parse Request", http.StatusBadRequest)
		return
	}

	id, status, err := h.Storage.Register(reqBody.Name, reqBody.Email, reqBody.Password)
	if !status || err != nil {
		http.Error(w, "Can't Register User", http.StatusBadRequest)
		return
	}

	response := models.UserResponse{
		Success: true,
		Message: "User Register Successfully",
		Data: struct {
			ID       int
			Email    string
			Username string
		}{ID: id, Email: reqBody.Email, Username: reqBody.Name},
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Can't Register User Response", http.StatusBadRequest)
		return
	}
}

func (h *UserHandlers) Authenticate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var reqBody models.AuthUserRequest
	json.NewDecoder(r.Body).Decode(&reqBody)

	user, status, err := h.Storage.Authenticate(reqBody.Email, reqBody.Password)
	if !status || err != nil {
		http.Error(w, "Can't Authenticate User", http.StatusBadRequest)
		return
	}

	token := token.CreateToken(user)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   24 * time.Now().Hour(),
		HttpOnly: true,
		Path:     "/",
		// Secure: true,
	}
	http.SetCookie(w, &cookie)
	response := models.UserResponse{
		Success: true,
		Message: "User Auth Successfully",
		Data: struct {
			ID       int
			Email    string
			Name string
		}{ID: user.ID, Email: reqBody.Email, Name : user.Name},
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Can't Auth User Response", http.StatusBadRequest)
		return
	}

}

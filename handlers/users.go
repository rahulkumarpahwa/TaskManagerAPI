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

	status, err := h.Storage.Register(reqBody.Name, reqBody.Email, reqBody.Password)
	if !status || err != nil {
		http.Error(w, "Can't Register User", http.StatusBadRequest)
		return
	}

	token := token.CreateToken(models.User{Username: reqBody.Name, Email: reqBody.Email, Password: reqBody.Password})

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
		Message: "User Register Successfully",
		Data: struct {
			Email    string
			Username string
		}{Email: reqBody.Email, Username: reqBody.Name},
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Can't Register User Response", http.StatusBadRequest)
		return
	}
}

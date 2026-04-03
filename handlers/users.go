package handlers

import (
	"TaskManager/data"
	"TaskManager/models"
	"encoding/json"
	"net/http"
	"time"
)

type UserHandlers struct {
	Storage data.UserRepositary
}

func (h *UserHandlers) Register(w http.ResponseWriter, r *http.Request) {

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


	


	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   24 * time.Now().Hour(),
		HttpOnly: true,
		Path:     "/",
		// Secure: true,
	}

	http.SetCookie(w, &cookie)

}

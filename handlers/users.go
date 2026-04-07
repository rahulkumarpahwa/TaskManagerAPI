package handlers

import (
	"TaskManager/data"
	"TaskManager/models"
	"TaskManager/token"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
			ID    int
			Email string
			Name  string
		}{ID: user.ID, Email: reqBody.Email, Name: user.Name},
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Can't Auth User Response", http.StatusBadRequest)
		return
	}

}

func (h *UserHandlers) AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || cookie == nil || cookie.Value == "" {
			log.Println("Cookie not found:", err)
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(cookie.Value,
			func(t *jwt.Token) (interface{}, error) {
				// Ensure the signing method is HMAC
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(token.GetJWTSecret()), nil
			},
		)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Get the id from claims
		id, ok := claims["id"].(int)
		if !ok {
			http.Error(w, "ID not found in token", http.StatusUnauthorized)
			return
		}

		// Inject email into the request context
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

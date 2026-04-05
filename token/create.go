package token

import (
	"TaskManager/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user models.User) string {
	jwtSecret := GetJWTSecret()

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Printf("Failed to Sign JWT %v", err)
		return ""
	}

	return tokenString
}

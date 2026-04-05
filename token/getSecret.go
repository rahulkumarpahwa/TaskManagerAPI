package token

import (
	"log"
	"os"
)

func GetJWTSecret() string {
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		jwtSecret = "default secret for the project"
		log.Print("JWT SECRET NOT SET, using the default value")
	}else {
		log.Print("Using the JWT SECRET")
	}
	return jwtSecret
}
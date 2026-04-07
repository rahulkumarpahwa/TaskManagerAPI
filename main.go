package main

import (
	"TaskManager/data"
	"TaskManager/handlers"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	server := http.NewServeMux()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnStr := os.Getenv("POSTGRES_CONNECTION")
	if dbConnStr == "" {
		log.Fatal("Error Getting the connection string")
	}

	db, err := sql.Open("pgx", dbConnStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	taskRepositary := &data.TaskRepositary{DB: db}

	taskHandler := handlers.TaskHandlers{Storage: taskRepositary}

	userRespositary := &data.UserRepositary{DB: db}

	userHandler := handlers.UserHandlers{Storage: userRespositary}

	server.HandleFunc("GET /", taskHandler.Health)
	server.HandleFunc("GET /task", taskHandler.GetTasks)
	server.Handle("POST /task", userHandler.AuthMiddleware(http.HandlerFunc(taskHandler.CreateTask)))
	server.Handle("PATCH /task/", userHandler.AuthMiddleware(http.HandlerFunc(taskHandler.UpdateTask)))
	server.Handle("DELETE /task/", userHandler.AuthMiddleware(http.HandlerFunc(taskHandler.DeleteTask)))

	server.HandleFunc("POST /user/auth/", userHandler.Authenticate)
	server.HandleFunc("POST /user/register/", userHandler.Register)

	http.ListenAndServe(":1999", server)
}

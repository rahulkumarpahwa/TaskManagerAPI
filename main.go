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
	server.HandleFunc("GET /", taskHandler.Health)
	server.HandleFunc("GET /task", taskHandler.GetTasks)
	server.HandleFunc("POST /task", taskHandler.CreateTask)
	server.HandleFunc("PATCH /task/", taskHandler.UpdateTask)
	server.HandleFunc("DELETE /task/", taskHandler.DeleteTask)

	userRespositary := &data.UserRepositary{DB: db}

	userHandler := handlers.UserHandlers{Storage: userRespositary}

	server.HandleFunc("POST /user", userHandler.Register)

	http.ListenAndServe(":1999", server)
}

package handlers

import (
	"TaskManager/data"
	"TaskManager/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type TaskHandlers struct {
	Storage data.TaskRepositaryModel
}

func (th *TaskHandlers) Health(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode("Health Check"); err != nil {
		http.Error(w, "Not able to get the data json", http.StatusBadRequest)
		return
	}
}

func (th *TaskHandlers) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := th.Storage.GetTasks()
	if err != nil {
		log.Printf("tasks can't be retrieved: %v", err)
		http.Error(w, "No Task Found!", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "Not able to get the data json", http.StatusBadRequest)
		return
	}
}

func (th *TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.CreateTask
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Not Able to Decode the Body", http.StatusBadRequest)
		return
	}
	var effectedRows = th.Storage.CreateTask(task)
	if effectedRows > 0 {
		w.Write([]byte("Task Added Successfully!"))
	} else {
		http.Error(w, "Task not added!", http.StatusBadRequest)
	}
}

func (th *TaskHandlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/task/"):]
	id_num, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id can't be found", http.StatusBadRequest)
		return
	}

	val := r.Context().Value("id")

	user_id, ok := val.(int)
	if !ok {
		http.Error(w, "Invalid or missing user ID", http.StatusUnauthorized)
		return
	}

	var task models.CreateTask
	w.Header().Set("Content-Type", "application/json")
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Not Able to Decode the Body", http.StatusBadRequest)
		return
	}

	updated_task, status, err := th.Storage.UpdateTask(id_num, task.Title, task.Description, task.Status, user_id, task.IsFavorite)
	if err != nil || !status {
		http.Error(w, "Not Able to Update the Task", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(updated_task); err != nil {
		http.Error(w, "Not Able to Decode the Updated Task", http.StatusBadRequest)
		return
	}
}

func (th *TaskHandlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Path[len("/task/"):]
	id_num, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id can't be found", http.StatusBadRequest)
	}

	val := r.Context().Value("id")

	user_id, ok := val.(int)
	if !ok {
		http.Error(w, "Invalid or missing user ID", http.StatusUnauthorized)
		return
	}

	deleted_task, status, err := th.Storage.DeleteTask(id_num, user_id)
	if err != nil || !status {
		http.Error(w, "Not Able to delete the Task : "+err.Error(), http.StatusBadRequest)
	}

	if status {
		response := models.UserResponse{
			Success: true,
			Message: "Task has been Deleted!",
			Data: struct {
				Message string
				Task    models.Task
			}{
				Message: "Deleted Task",
				Task:    deleted_task,
			},
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Not Able to Decode the Deleted Task", http.StatusBadRequest)
			return
		}
	}
}

func (h *TaskHandlers) SetFavoriteTask(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Path[len("/task/set-favorite/"):]
	taskIdNum, err := strconv.Atoi(idStr)
	fmt.Print(taskIdNum)
	
	if err != nil {
		http.Error(w, "Task Id can't be found", http.StatusBadRequest)
		return
	}

	val := r.Context().Value("id")

	user_id, ok := val.(int)
	if !ok {
		http.Error(w, "Invalid or missing user ID", http.StatusUnauthorized)
		return
	}

	status, err := h.Storage.SetFavoriteTask(taskIdNum, user_id)

	response := models.UserResponse{
		Success: true,
		Message: "Task has been Favorited!",
		Data:    status,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Not Able to Encode the Favorite Task", http.StatusBadRequest)
		return
	}
}

func (h *TaskHandlers) GetFavoriteTasks(w http.ResponseWriter, r *http.Request) {

	val := r.Context().Value("id")

	user_id, ok := val.(int)
	if !ok {
		http.Error(w, "Invalid or missing user ID", http.StatusUnauthorized)
		return
	}

	fav_tasks, err := h.Storage.GetFavoriteTasks(user_id)
	if err != nil {
		http.Error(w, "Not Able to get the favorite tasks", http.StatusBadRequest)
		return
	}

	response := models.UserResponse{
		Success: true,
		Message: "Task has been Favorited!",
		Data: struct {
			Message string
			Tasks   []models.Task
		}{
			Message: "Favorites",
			Tasks:   fav_tasks,
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Not Able to Encode the Favorite Task", http.StatusBadRequest)
		return
	}
}

package handlers

import (
	"TaskManager/data"
	"TaskManager/models"
	"encoding/json"
	"net/http"
)

type TaskHandlers struct {
	Storage data.TaskRepositaryModel
}

func (th *TaskHandlers) GetTasks(w http.ResponseWriter, r *http.Request) {

}

func (th *TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var tasks models.CreateTask
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, "Not Able to Decode the Body", http.StatusBadRequest)
		return
	}
}

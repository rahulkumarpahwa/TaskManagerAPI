package handlers

import (
	"TaskManager/data"
	"TaskManager/models"
	"encoding/json"
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
	tasks := th.Storage.GetTasks()
	if tasks == nil {
		log.Fatal("Tasks can't be found!")
		http.Error(w, "No Task Found!", http.StatusBadRequest)
	}
	err := json.NewEncoder(w).Encode(tasks)
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
	idStr := r.URL.Path[len("/tasks/"):]
	var task models.CreateTask
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Not Able to Decode the Body", http.StatusBadRequest)
		return
	}
	id_num, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("Id can't be found %v", err)
		http.Error(w, "Id can't be found", http.StatusBadRequest)
	}
	th.Storage.UpdateTask(id_num, task.Title, task.Description, task.Status)

}

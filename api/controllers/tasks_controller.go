package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/LucasStarlingdePaulaSalles/insprTasks/api/models"
	"github.com/LucasStarlingdePaulaSalles/insprTasks/api/responses"
)

// CreateTask is a handler function for creating new tasks
func (server *Server) CreateTask(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	task := models.Task{}
	err = json.Unmarshal(body, &task)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	task.Prepare()
	err = task.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	taskCreated, err := task.SaveTask(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, taskCreated.ID))
	responses.JSON(w, http.StatusCreated, taskCreated)
}

//GetTasks is a handler function for getting all tasks, unfiltered
func (server *Server) GetTasks(w http.ResponseWriter, r *http.Request) {

	task := models.Task{}
	tasks, err := task.FindAllTasks(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, tasks)
}
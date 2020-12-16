package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LucasStarlingdePaulaSalles/insprTasks/frontend/models"
)

func GetAllTasks() {
	resp, err := http.Get("http://localhost:8080/tasks")
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}
	defer resp.Body.Close()
	tasks := []models.Task{}
	json.NewDecoder(resp.Body).Decode(&tasks)
	PrintTasks(tasks...)
}

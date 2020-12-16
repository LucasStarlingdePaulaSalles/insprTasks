package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LucasStarlingdePaulaSalles/insprTasks/frontend/models"
)

func Work() {
	var ID int
	fmt.Print("Task ID: ")
	fmt.Scan(&ID)
	url := "http://localhost:8080/work/" + strconv.Itoa(ID)
	bites := []byte(`{}`)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bites))
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}
	defer req.Body.Close()
	tasks := models.Task{}
	json.NewDecoder(resp.Body).Decode(&tasks)
	PrintTasks(tasks)
}

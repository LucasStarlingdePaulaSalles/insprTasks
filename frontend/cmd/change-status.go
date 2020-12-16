package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/LucasStarlingdePaulaSalles/insprTasks/frontend/models"
)

//ChangeStatus sends a requisition to change a task's status
func ChangeStatus() {
	var ID int
	var status uint8
	fmt.Print("Task ID: ")
	_, err := fmt.Scan(&ID)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	url := "http://localhost:8080/status/" + strconv.Itoa(ID)
	status = getStatus()
	if status == 3 {
		if !closingTaskConfirm() {
			return
		}
	}
	newStatus := models.NewStatusDTO{}
	newStatus.NewStatus = status
	body, err := json.Marshal(newStatus)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
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

func getStatus() uint8 {
	var statusCode uint8
while:
	for {
		fmt.Println("Choose a status by it's code:")
		w := tabwriter.NewWriter(os.Stdout, 8, 0, 1, ' ', tabwriter.Debug)
		fmt.Fprintf(w, "Status\t Code\n")
		fmt.Fprintf(w, "ToDo\t 0\n")
		fmt.Fprintf(w, "Working\t 1\n")
		fmt.Fprintf(w, "Done\t 2\n")
		fmt.Fprintf(w, "Closed\t 3\n")
		w.Flush()
		fmt.Print("Code for new status: ")
		_, err := fmt.Scan(&statusCode)
		if err == nil && statusCode >= 0 && statusCode <= 3 {
			break while
		}
	}
	return statusCode
}

//CloseTask sends a requisition to close a task
func CloseTask() {
	var ID int
	fmt.Print("Task ID: ")
	_, err := fmt.Scan(&ID)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	url := "http://localhost:8080/status/" + strconv.Itoa(ID)
	newStatus := models.NewStatusDTO{}
	newStatus.NewStatus = 3
	if !closingTaskConfirm() {
		return
	}
	body, err := json.Marshal(newStatus)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
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

func closingTaskConfirm() bool {
	fmt.Println("Closing a task cannot be undone!")
	fmt.Println("Continue anyway? (yes/no)")
	var input string
	fmt.Scan(&input)
	if input == "yes" || input == "y" {
		return true
	}
	return false
}

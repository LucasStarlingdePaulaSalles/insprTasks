package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/LucasStarlingdePaulaSalles/insprTasks/myTasks/models"
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

func GetFilteredTasks() {
	var input string
while:
	for {
		fmt.Println("Tasks can be filtered multiple ways:")
		w := tabwriter.NewWriter(os.Stdout, 12, 0, 2, ' ', tabwriter.Debug)
		fmt.Fprintf(w, "Command\t Filter\n")
		fmt.Fprintf(w, "priority\t Filter by chozen level of priority\n")
		fmt.Fprintf(w, "deadline\t Filter by chozen deadline date\n")
		fmt.Fprintf(w, "added\t Filter by chozen date of creation\n")
		fmt.Fprintf(w, "ID\t Filter by chozen ID (if exists)\n")
		fmt.Fprintf(w, "status\t Filter by chozen status\n")
		w.Flush()
		fmt.Print("Filter type: ")
		_, err := fmt.Scan(&input)
		if err == nil {
			switch input {
			case "priority":
				getTaskByValues("priority")
			case "deadline":
				getTaskByDate("deadline")
			case "added":
				getTaskByDate("creation")
			case "ID":
				getTaskByValues("ID")
			case "status":
				getTaskByValues("status")
			}
			break while
		}
	}
}

//GetCurrentTask querrys the backend for the tasr on 'Working' status
func GetCurrentTask() {
	working := 1
	filter := models.NumericFilterDTO{}
	filter.Field = "status"
	filter.Value = working

	body, err := json.Marshal(filter)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	req, err := http.NewRequest("GET", "http://localhost:8080/tasks/value", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}
	defer req.Body.Close()
	tasks := []models.Task{}
	json.NewDecoder(resp.Body).Decode(&tasks)
	fmt.Println("Currently working on")
	PrintTasks(tasks...)
}

func getTaskByDate(param string) {
	var dateStr string
	fmt.Printf("Date (%s) [dd/mm/yyyy]: ", param)
	fmt.Scan(&dateStr)
	date, err := simpeDateParse(dateStr)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}

	filter := models.DateFilterDTO{}
	filter.Field = param
	filter.Year = date.Year()
	filter.Day = date.Day()
	filter.Month = int(date.Month())
	body, err := json.Marshal(filter)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	req, err := http.NewRequest("GET", "http://localhost:8080/tasks/date", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}
	defer req.Body.Close()
	tasks := []models.Task{}
	json.NewDecoder(resp.Body).Decode(&tasks)
	PrintTasks(tasks...)
}

func getTaskByValues(param string) {
	var value int
	fmt.Printf("Value (%s) ", param)
	fmt.Scan(&value)

	filter := models.NumericFilterDTO{}
	filter.Field = param
	filter.Value = value

	body, err := json.Marshal(filter)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	req, err := http.NewRequest("GET", "http://localhost:8080/tasks/value", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}
	defer req.Body.Close()
	tasks := []models.Task{}
	json.NewDecoder(resp.Body).Decode(&tasks)
	PrintTasks(tasks...)
}

func simpeDateParse(dateStr string) (time.Time, error) {
	layout := "02/01/2006"
	return time.Parse(layout, dateStr)
}

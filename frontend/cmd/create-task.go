package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/LucasStarlingdePaulaSalles/insprTasks/frontend/models"
)

//CreateTask receives user inputs for creating a task
func CreateTask() {
	fmt.Println("Creating a task...\n")
	task := models.Task{}
	fmt.Print("Task name: ")
	var err error
	var text string
	//Title input
	text, err = lineReader()
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	textParse(&task.Title, text)
	//Description input
	fmt.Print("Task description (up to 255 char): ")
	text, err = lineReader()
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	textParse(&task.Description, text)
	//Priority input
	var num int
	fmt.Print("Task priority [1 : 10]: ")
	_, err = fmt.Scan(&num)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	priorityParse(&task, num)
	//Deadline input
	fmt.Print("Task deadline [dd/mm/yyyy]: ")
	fmt.Scan(&text)
	err = dateParse(&task, text)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	//TimeEstimate input
	fmt.Print("Task time estimate (h): ")
	_, err = fmt.Scan(&num)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	estimateParse(&task, num)
	//Dependencies input
	getDependencies(&task)
	fmt.Println()
	fmt.Println("Your task:")
	PrintTasks(task)
	fmt.Println("Create? (yes/no)")
	fmt.Scan(&text)
	fmt.Println(text)
	if text == "yes" || text == "y" {
		sendTask(task)
	} else {
		return
	}
}

func sendTask(task models.Task) {
	body, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/task", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro: %s \n", err)
	}

	defer resp.Body.Close()
	if resp.Status == "201 Created" {
		fmt.Println("Created!")
	} else {
		fmt.Println("Something went wrong:")
		fmt.Println(resp.Status)
		respBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Error:", string(respBody))
	}
}

func lineReader() (string, error) {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	line = line[:len(line)-1]
	return line, err
}

func textParse(field *string, value string) {
	if len(value) > 256 {
		*field = value[0:255]
	} else {
		*field = value
	}
}

func dateParse(task *models.Task, dateStr string) error {
	layout := "02/01/2006"
	var err error
	task.Deadline, err = time.Parse(layout, dateStr)
	return err
}

func priorityParse(task *models.Task, priority int) {
	if priority < 1 {
		priority = 1
	} else if priority > 10 {
		priority = 10
	}
	task.Priority = uint8(priority)
}

func estimateParse(task *models.Task, duration int) {
	task.TimeEstimate = time.Duration(duration) * time.Hour
}

func getDependencies(task *models.Task) {
	var ID int
	var input string
	task.Dependencies = ""
	fmt.Println("Is this task dependant on any tasks? (yes/no)")
	fmt.Scan(&input)
	if input == "yes" || input == "y" {
		fmt.Print("Blocker ID: ")
		fmt.Scan(&ID)
		task.Dependencies = strconv.Itoa(ID)
	} else {
		return
	}

while:
	for {
		fmt.Println("Any other? (yes/no)")
		fmt.Scan(&input)
		if input == "yes" || input == "y" {
			fmt.Print("Blocker ID: ")
			fmt.Scan(&ID)
			task.Dependencies = task.Dependencies + ";" + strconv.Itoa(ID)
		} else {
			break while
		}
	}
}

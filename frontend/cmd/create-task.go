package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
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

	getTitle(&task)
	getDescription(&task)
	getPriority(&task)
	getDeadline(&task)
	getTimeEstimate(&task)
	getDependencies(&task)

	fmt.Println()
	fmt.Println("Your task:")
	PrintTasks(task)

	fmt.Println("Create? (yes/no)")
	var input string
	fmt.Scan(&input)
	fmt.Println(input)
	if input == "yes" || input == "y" {
		sendTask(task)
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

func getTitle(task *models.Task) {
	for {
		fmt.Print("Task name: ")
		text, err := lineReader()
		if err == nil {
			if text != "" {
				textParse(&task.Title, text)
				return
			}
			err = errors.New("Title is required")
		}
		fmt.Printf("Erro: %s \n\n", err)
	}
}

func getDescription(task *models.Task) {
	for {
		fmt.Print("Task description (up to 255 char): ")
		text, err := lineReader()
		if err == nil {
			textParse(&task.Description, text)
			return
		}
		fmt.Printf("Erro: %s \n", err)
	}
}

func getPriority(task *models.Task) {
	var priority int
	for {
		fmt.Print("Task priority [1 : 10]: ")
		_, err := fmt.Scan(&priority)
		if err == nil {
			if priority > 0 && priority < 11 {
				task.Priority = uint8(priority)
				return
			}
			err = errors.New("Priority must be an integer between 1 and 10")
		}
		fmt.Printf("Erro: %s \n", err)
	}
}

func getDeadline(task *models.Task) {
	layout := "02/01/2006"
	var err error
	var input string
	for {
		fmt.Print("Task deadline [dd/mm/yyyy]: ")
		fmt.Scan(&input)
		task.Deadline, err = time.Parse(layout, input)
		if err == nil {
			return
		}
		err = errors.New("Connot parse '" + input + "' as 'dd/mm/yyyy'")
		fmt.Printf("Erro: %s \n", err)
	}
}

func getTimeEstimate(task *models.Task) {
	var input string
	var err error
	var hours float64
	for {
		fmt.Print("Task time estimate (h): ")
		_, err = fmt.Scan(&input)
		if err == nil {
			hours, err = strconv.ParseFloat(input, 64)
			fmt.Println(err)
			if err == nil {
				task.TimeEstimate = time.Duration(hours*60) * time.Minute
				return
			}
		}
		fmt.Printf("Erro: %s \n", err)
	}
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

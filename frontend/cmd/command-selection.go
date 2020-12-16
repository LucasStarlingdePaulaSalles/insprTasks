package cmd

import (
	"fmt"
)

func Init() {
	Help()
	Selector()
}

func Help() {
	fmt.Println("Command options:")
	fmt.Println("help   |                 Show this information")
	fmt.Println("exit   |                     Close application")
	fmt.Println("new    |                     Create a new task")
	fmt.Println("all    |                        Show all tasks")
	fmt.Println("work   |               Start working on a task")
	fmt.Println("stop   |    Stop work on currently active task")
	fmt.Println()
}

func Selector() {
	var input string
while:
	for {
		fmt.Print("->")
		fmt.Scan(&input)
		fmt.Println()
		fmt.Println()

		switch input {
		case "new":
			CreateTask()
		case "help":
			Help()
		case "all":
			GetAllTasks()
		case "work":
			Work()
		case "stop":
			StopWork()
		case "exit":
			break while
		}
	}

}

package cmd

import (
	"fmt"
)

//Init starts the program on interactive mode, on this mode the user can execute multiple commands.
//Interactive mode is only interrupted by the user.
func Init() {
	help()
	var input string
while:
	for {
		fmt.Print("->")
		fmt.Scan(&input)
		fmt.Println()
		fmt.Println()
		exit := Selector(input)
		if exit {
			break while
		}
	}
}

//Selector executes a function based on its command
func Selector(arg string) bool {
	switch arg {
	case "all":
		GetAllTasks()
	case "current":
		GetCurrentTask()
	case "new":
		CreateTask()
	case "help":
		help()
	case "work":
		Work()
	case "stop":
		StopWork()
	case "status":
		ChangeStatus()
	case "close":
		CloseTask()
	case "tasks":
		GetFilteredTasks()
	case "exit":
		fmt.Println("Goodbye!")
		return true
	default:
		fmt.Println("Unknown command :(")
		fmt.Println("Use 'help' for details")
		fmt.Println()
	}
	return false
}

func help() {
	fmt.Println("Command options:")
	fmt.Println("help      |                Show's this information")
	fmt.Println("exit      |                      Close application")
	fmt.Println("new       |                      Create a new task")
	fmt.Println("current   |   Show's task currently been worked on")
	fmt.Println("all       |                       Show's all tasks")
	fmt.Println("work      |                Start working on a task")
	fmt.Println("stop      |     Stop work on currently active task")
	fmt.Println("status    |                 Change a task's status")
	fmt.Println("close     |                           Close a task")
	fmt.Println("tasks     |       Show's tasks that match a filter")
	fmt.Println()
}

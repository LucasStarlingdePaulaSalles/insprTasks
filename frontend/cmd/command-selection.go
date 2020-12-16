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
	fmt.Println("help   |   Show this information")
	fmt.Println("exit   |       Close application")
	fmt.Println("new    |       Create a new task")
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
		case "exit":
			break while
		}
	}

}

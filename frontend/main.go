package main

import (
	"fmt"
	"os"

	"github.com/LucasStarlingdePaulaSalles/insprTasks/frontend/cmd"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		cmd.Selector(args[1])
	} else if len(args) == 1 {
		fmt.Println("Welcome to InsprTasks!")
		cmd.Init()
	} else {
		fmt.Println("Error: Incorrect argument structure")
	}

}

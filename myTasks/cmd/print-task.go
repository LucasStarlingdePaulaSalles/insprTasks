package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/LucasStarlingdePaulaSalles/insprTasks/myTasks/models"
)

//PrintTasks prints a task or list of tasks on standard table format
func PrintTasks(tasks ...models.Task) {
	const padding = 3

	w := tabwriter.NewWriter(os.Stdout, 12, 0, padding, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "ID\t Name\t Status\t Priority\t Deadline\t TimeEstimate\t Work done\t Dependencies\t Description\n")
	for _, task := range tasks {
		fmt.Fprintf(w, "%d\t %s\t %s\t %d\t %s\t %.1f\t %.1f\t %s\t %s\n",
			task.ID, task.Title, parseStastus(task.Status), task.Priority, task.Deadline.Format("02/01/2006"),
			task.TimeEstimate.Hours(), task.WorkedFor.Hours(), task.Dependencies, task.Description)
	}
	w.Flush()
}

func parseStastus(status uint8) string {
	switch status {
	case 0:
		return "ToDo"
	case 1:
		return "Working"
	case 2:
		return "Done"
	default:
		return "Closed"
	}
}

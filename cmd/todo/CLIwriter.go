package todo

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func tabwriteFull(tasks []Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprint(w, "ID\tTask\tCreated\tStatus\n")
	for _, task := range tasks {
		fmt.Fprint(w, task.ID, "\t", task.Task, "\t", task.Created, "\t", task.Status, "\n")
	}
	w.Flush()
}

func tabwriteTrunc(tasks []Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprint(w, "ID\tTask\tCreated\n")

	for _, task := range tasks {
		if task.Status == "undone" {
			fmt.Fprint(w, task.ID, "\t", task.Task, "\t", task.Created,"\n")
		}
	}
	w.Flush()
}

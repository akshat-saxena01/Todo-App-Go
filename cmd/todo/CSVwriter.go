package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func writeNew(file *os.File) {
	wCsv := csv.NewWriter(file)

	wCsv.Write([]string{"ID", "Task", "Created", "Status"})
	wCsv.Flush()
}

func writeCsv(tasks []Task) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()

	wCsv := csv.NewWriter(file)

	wCsv.Write([]string{"ID", "Task", "Created", "Status"})
	for _, task := range tasks {
		wCsv.Write([]string{strconv.Itoa(task.ID), task.Task, task.Created, task.Status})
	}

	wCsv.Flush()
}

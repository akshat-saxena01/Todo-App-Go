package todo

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func StructComplete(reader *csv.Reader, TaskId int) []Task {
	var tasks []Task
	_, err := reader.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for counter := 1; ; counter++ {
		row, err := reader.Read()
		if err == io.EOF {
			if TaskId >= counter {
				fmt.Fprintf(os.Stderr, "Task ID %d does not exist\n", TaskId)
			}
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		idint, err := strconv.Atoi(row[id])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if idint == TaskId {
			flag := false
			if row[status] == "undone" {
				flag = true
			}
			task := Task{
				ID:      idint,
				Task:    row[task],
				Created: row[created],
				Status:  "done",
			}
			if flag {
				fmt.Printf("Task %s is done\n", task.Task)
			}
			tasks = append(tasks, task)
			continue
		}
		task := Task{
			ID:      idint,
			Task:    row[task],
			Created: row[created],
			Status:  row[status],
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func CompleteTask() {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var tasks []Task

	if len(os.Args) == 3 {
		TaskId, err := strconv.Atoi(os.Args[2])
		if err == strconv.ErrSyntax {
			fmt.Fprintln(os.Stderr, err, "Please enter a valid number")
			os.Exit(1)
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		tasks = StructComplete(reader, TaskId)
	} else {
		fmt.Fprintln(os.Stderr, "Invalid Input")
		os.Exit(1)
	}

	if len(tasks) >= 1 {
		writeCsv(tasks)
	} else {
		fmt.Fprint(os.Stderr, "No tasks exist\n")
	}

}

package todo

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func StructDelete(reader *csv.Reader, TaskId int) []Task {
	var tasks []Task
	_, err := reader.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	newCount := 1
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
			continue
		}
		task := Task{
			ID:      newCount,
			Task:    row[task],
			Created: row[created],
			Status:  row[status],
		}
		tasks = append(tasks, task)
		newCount++
	}
	return tasks
}

func DeleteAll() {
	file, _ := os.Create(path)
	writeNew(file)
}

func DeleteTask() {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var tasks []Task

	if len(os.Args) == 3 && os.Args[2] == "--all" {
		DeleteAll()
	} else if len(os.Args) == 3 {
		TaskId, err := strconv.Atoi(os.Args[2])
		if err == strconv.ErrSyntax {
			fmt.Fprintln(os.Stderr, err, "Please enter a valid number")
			os.Exit(1)
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		tasks = StructDelete(reader, TaskId)
	} else {
		fmt.Fprintln(os.Stderr, "Invalid Input")
		os.Exit(1)
	}

	writeCsv(tasks)
	if tasks == nil {
		fmt.Fprintln(os.Stderr, "No tasks exist")
	}
}

package todo

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func StructSlice(reader *csv.Reader) []Task {
	var tasks []Task
	_, err := reader.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		idint, err := strconv.Atoi(row[id])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
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

func ListTask() {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var tasks []Task = StructSlice(reader)

	if len(os.Args) == 2 {
		tabwriteTrunc(tasks)
	} else if len(os.Args) == 3 && os.Args[2] == "--all" {
		tabwriteFull(tasks)
	} else {
		fmt.Fprintln(os.Stderr, "Invalid Input")
		os.Exit(1)
	}
}

package todo

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func StructAdd(reader *csv.Reader, s []string) []Task {
	var tasks []Task
	_, err := reader.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	lastIndex := 0
	for ; ; lastIndex++ {
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
	lastIndex++
	currentTime := time.Now()
	time := fmt.Sprintf("%s %d, %d, %d:%d", currentTime.Month(), currentTime.Day(), currentTime.Year(), currentTime.Hour(), currentTime.Minute())
	tasks = append(tasks, Task{lastIndex, strings.Join(s, " "), time, "undone"})

	return tasks
}

func AddTask() {
	_, err := os.Open(path)
	if os.IsNotExist(err) {
		file, _ := os.Create(path)
		writeNew(file)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var tasks []Task

	if len(os.Args) >= 3 {
		tasks = StructAdd(reader, os.Args[2:])
	} else {
		fmt.Fprintln(os.Stderr, "Invalid Input")
		os.Exit(1)
	}
	writeCsv(tasks)
}

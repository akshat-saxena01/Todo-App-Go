package todo

import (
	"fmt"
	"os"
)

var PathDir string = func() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	s := fmt.Sprintf("%s/Documents/gtask", home)
	return s
}()

var path string = fmt.Sprintf("%s/task.csv", PathDir)

type Task struct {
	ID      int
	Task    string
	Created string
	Status  string
}

const (
	id = iota
	task
	created
	status
)

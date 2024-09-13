package main

import (
	"fmt"
	"gtask/todo"
	"os"
)

func main() {
	// Check if os.Args[1] exists
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Too few arguments")
		os.Exit(1)
	}

	_,err :=os.Stat(todo.PathDir)
	if os.IsNotExist(err) {
		os.MkdirAll(todo.PathDir,os.ModePerm)
		fmt.Println("Gtask directory created at "+todo.PathDir)
	} else if err != nil {
		fmt.Fprintln(os.Stderr,err)
	}

	switch os.Args[1] {
	case "add":
		todo.AddTask()
	case "list":
		todo.ListTask()
	case "complete":
		todo.CompleteTask()
	case "delete":
		todo.DeleteTask()
	default:
		fmt.Fprintln(os.Stderr, "Invalid Input, try again")
		os.Exit(1)
	}
}

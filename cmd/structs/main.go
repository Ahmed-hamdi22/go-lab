package main

import "fmt"

type LabExperiment struct {
	Title    string
	Category string
	Done     bool
}

func (e LabExperiment) Summary() string {
	status := "In Progress"
	if e.Done {
		status = "Completed"
	}
	return fmt.Sprintf("[%s] %s - %s", e.Category, e.Title, status)
}

func main() {
	fmt.Println("=== Structs Lab ===")

	exp := LabExperiment{
		Title:    "Go Fundamentals & Setup",
		Category: "Basics",
		Done:     true,
	}

	fmt.Println(exp.Summary())
}

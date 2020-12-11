package common

import "github.com/pellared/taskflow"

// TaskFmt runs go fmt.
func TaskFmt() taskflow.Task {
	return taskflow.Task{
		Name:        "fmt",
		Description: "go fmt",
		Command:     taskflow.Exec("go", "fmt", "./..."),
	}
}

package main

import (
	"os"
	"path/filepath"

	"github.com/pellared/taskflow"
)

func main() {
	tasks := &taskflow.Taskflow{}

	clean := tasks.MustRegister(taskflow.Task{
		Name:        "clean",
		Description: "remove files created during build",
		Command:     taskClean,
	})

	fmt := tasks.MustRegister(taskflow.Task{
		Name:        "fmt",
		Description: "go fmt",
		Command:     taskFmt,
	})

	test := tasks.MustRegister(taskflow.Task{
		Name:        "test",
		Description: "go test with race detector and code covarage",
		Command:     taskTest,
	})

	tasks.MustRegister(taskflow.Task{
		Name:        "all",
		Description: "build pipeline",
		Dependencies: taskflow.Deps{
			clean,
			fmt,
			test,
		},
	})

	tasks.Main()
}

func taskClean(tf *taskflow.TF) {
	files, err := filepath.Glob("coverage.*")
	if err != nil {
		tf.Fatalf("glob failed: %v", err)
	}
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			tf.Errorf("failed to remove %s: %v", file, err)
			continue
		}
		tf.Logf("removed %s", file)
	}
}

func taskFmt(tf *taskflow.TF) {
	if err := tf.Exec("", nil, "go", "fmt", "./..."); err != nil {
		tf.Errorf("go fmt: %v", err)
	}
}

func taskTest(tf *taskflow.TF) {
	if err := tf.Exec("", nil, "go", "test", "-race", "-covermode=atomic", "-coverprofile=coverage.out", "./..."); err != nil {
		tf.Errorf("go test: %v", err)
	}
	if err := tf.Exec("", nil, "go", "tool", "cover", "-html=coverage.out", "-o", "coverage.html"); err != nil {
		tf.Errorf("go tool cover: %v", err)
	}
}

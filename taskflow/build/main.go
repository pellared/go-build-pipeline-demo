package main

import (
	"os"
	"path/filepath"

	"github.com/pellared/taskflow"

	"github.com/pellared/go-build-pipeline-demo/taskflow/common"
)

func main() {
	tasks := &taskflow.Taskflow{}

	clean := tasks.MustRegister(taskClean())
	fmt := tasks.MustRegister(common.TaskFmt())
	test := tasks.MustRegister(taskTest())

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

func taskClean() taskflow.Task {
	return taskflow.Task{
		Name:        "clean",
		Description: "remove files created during build",
		Command: func(tf *taskflow.TF) {
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
		},
	}
}

func taskTest() taskflow.Task {
	return taskflow.Task{
		Name:        "test",
		Description: "go test with race detector and code covarage",
		Command: func(tf *taskflow.TF) {
			if err := tf.Cmd("go", "test", "-race", "-covermode=atomic", "-coverprofile=coverage.out", "./...").Run(); err != nil {
				tf.Errorf("go test: %v", err)
			}
			if err := tf.Cmd("go", "tool", "cover", "-html=coverage.out", "-o", "coverage.html").Run(); err != nil {
				tf.Errorf("go tool cover: %v", err)
			}
		},
	}
}

package main

import (
	"os"
	"path/filepath"

	"github.com/pellared/taskflow"

	"github.com/pellared/go-build-pipeline-demo/taskflow/common"
)

func main() {
	tasks := &taskflow.Taskflow{}

	c := common.Register(tasks)

	clean := tasks.MustRegister(taskflow.Task{
		Name:        "clean",
		Description: "remove files created during build",
		Command:     taskClean,
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
			c.Fmt,
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

func taskTest(tf *taskflow.TF) {
	if err := tf.Exec("", nil, "go", "test", "-race", "-covermode=atomic", "-coverprofile=coverage.out", "./..."); err != nil {
		tf.Errorf("go test: %v", err)
	}
	if err := tf.Exec("", nil, "go", "tool", "cover", "-html=coverage.out", "-o", "coverage.html"); err != nil {
		tf.Errorf("go tool cover: %v", err)
	}
}

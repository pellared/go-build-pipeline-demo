package common

import "github.com/pellared/taskflow"

// Tasks contains common registered tasks.
type Tasks struct {
	Fmt taskflow.RegisteredTask
}

// Register registers common tasks and returns them so they can be used as dependencies.
func Register(tasks *taskflow.Taskflow) Tasks {
	fmt := tasks.MustRegister(taskflow.Task{
		Name:        "fmt",
		Description: "go fmt",
		Command:     taskFmt,
	})

	return Tasks{
		Fmt: fmt,
	}
}

func taskFmt(tf *taskflow.TF) {
	if err := tf.Exec("", nil, "go", "fmt", "./..."); err != nil {
		tf.Errorf("go fmt: %v", err)
	}
}

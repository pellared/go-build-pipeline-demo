package main

import (
	"testing"

	"github.com/pellared/taskflow"
)

func Test(t *testing.T) {
	task := taskClean()
	r := taskflow.Runner{}
	if got := r.Run(task.Command); got.Failed() {
		t.FailNow()
	}
}

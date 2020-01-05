// +build !windows

package action

import (
	"os/exec"
)

func (w *worker) workerPath() string {
	return "../bin/worker"
}

func (w *worker) setFlag(cmd *exec.Cmd) {
}

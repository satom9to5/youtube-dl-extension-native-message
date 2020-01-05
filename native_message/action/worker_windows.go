package action

import (
	"os/exec"
	"syscall"
)

// https://docs.microsoft.com/en-us/windows/win32/procthread/process-creation-flags
const (
	CREATE_BREAKAWAY_FROM_JOB = 0x01000000
	CREATE_NEW_CONSOLE        = 0x00000010
)

func (w *worker) workerPath() string {
	return "../bin/worker.exe"
}

func (w *worker) setFlag(cmd *exec.Cmd) {
	if w.Browser == "Firefox" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			CreationFlags: CREATE_BREAKAWAY_FROM_JOB | CREATE_NEW_CONSOLE | syscall.CREATE_NEW_PROCESS_GROUP,
		}
	}
}

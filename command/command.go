package command

import (
	"os/exec"
	"syscall"

	"github.com/swapbyt3s/lightflow/common"
)

func Execute(cmd string) (stdout string, exitcode int) {
	out, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitcode = ws.ExitStatus()
		}
	}
	stdout = string(out[:])
	stdout = common.TrimNewlines(stdout)

	return
}

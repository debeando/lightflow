package execute

import (
	"os/exec"
	"syscall"

	"github.com/debeando/lightflow/common"
)

func Execute(cmd string, dryrun bool) (stdout string, exitcode int) {
	// var dryrun_arg string

	if dryrun {
		// dryrun_arg = "-n"
	}

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

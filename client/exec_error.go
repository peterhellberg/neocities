package client

import (
	"os/exec"
	"syscall"
)

// ExecError contains one error and an exit code.
type ExecError struct {
	Err      error
	ExitCode int
}

func (execError *ExecError) Error() string {
	return execError.Err.Error()
}

func newExecError(err error) ExecError {
	exitCode := 0
	if err != nil {
		exitCode = 1
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				exitCode = status.ExitStatus()
			}
		}
	}

	return ExecError{Err: err, ExitCode: exitCode}
}

package daemonize

import (
	"os"
	"os/exec"
	"errors"
)

func KillByPid(pid int) errors.Error {
	killcmd := exec.Command("kill", string(pid))
	killcmd.Stdout = os.Stdout
	err := killcmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

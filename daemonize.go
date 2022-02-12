package daemonize

import (
	"os"
	"os/exec"
)

func KillByPid(pid int) error {
	killcmd := exec.Command("kill", string(pid))
	killcmd.Stdout = os.Stdout
	err := killcmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

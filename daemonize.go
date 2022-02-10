package daemonize

import (
	"log"
	"os"
	"os/exec"
	"errors"
)

func Cmd(inp exec.Cmd) errors.Error {
	inp.Stdout = os.Stdout
	err := inp.Start()
	if err != nil {
		return err
	}
	log.Printf("Background Process %d started", inp.Process.Pid)
	return nil
}

func Kill(pid int) errors.Error {
	killcmd := exec.Command("kill", string(pid))
	killcmd.Stdout = os.Stdout
	err := killcmd.Wait()
	if err != nil {
		return err
	}
	log.Printf("Process %d killed", inp.Process.Pid)
	return nil
}

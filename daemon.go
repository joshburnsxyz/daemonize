package daemonize

import (
	"os"
	"os/exec"
)

type Daemon struct {
	Cmd *exec.Cmd
}

func New(inp *exec.Cmd) *Daemon {
	inp.Stdout = os.Stdout
	return &Daemon{Cmd: inp}
}

func (d *Daemon) Start() error {
	d.Cmd.Stdout = os.Stdout
	err := d.Cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func (d *Daemon) Kill() error {
	err := d.Cmd.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}

package daemonize

type Daemon struct {
	Cmd exec.Cmd
}

func New(inp exec.Cmd) *Daemon {
	inp.Stdout = os.Stdout
	return nil, &Daemon{Cmd: inp}
}

func (d *Daemon) Start() errors.Error {
	d.Cmd.Stdout = os.Stdout
	err := d.Cmd.Start()
	if err != nil {
		return err
	}
	return nil
}
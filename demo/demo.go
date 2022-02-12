package main

import (
	"os/exec"
	"log"
	"github.com/joshburnsxyz/daemonize"
)

func main() {
	cmd := exec.Command("sleep", "5")
	daemon := daemonize.New(cmd)
	err := daemon.Start()
	if err != nil {
		log.Fatalf("Daemon Error %s\n", err.Error())
	}
}

# Daemonize
Launch and manage daemons in Go

## Installation

```console
$ go get github.com/joshburnsxyz/go-daemonize
```

## Usage

__TLDR Example__

```go
package main

import (
  "os/exec"
  "log"
  "github.com/joshburnsxyz/go-daemonize"
)

func main() {
  cmd := exec.Command("python", "./some-script.py")
  daemon := daemonize.New(cmd)
  err := daemon.Start()
  if err != nil {
    log.Fatalf("Daemon Error %s\n", err.Error())
  }
}
```

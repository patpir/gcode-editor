package main

import (
	"os/signal"
	"syscall"

	"github.com/patpir/gcode-pattern-detect/commands"
)

func main() {
	signal.Ignore(syscall.SIGPIPE)

	commands.Execute()
}


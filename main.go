package main

import (
	"os/signal"
	"syscall"

	"github.com/patpir/gcode-editor/commands"
)

func main() {
	signal.Ignore(syscall.SIGPIPE)

	commands.Execute()
}


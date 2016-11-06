package main

import (
	"os"

	"github.com/spiegel-im-spiegel/pi/cmd"
)

func main() {
	cmd.Execute()
	os.Exit(cmd.ExitCode)
}

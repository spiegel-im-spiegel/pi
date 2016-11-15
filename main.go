/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */
package main

import (
	"os"

	"github.com/spiegel-im-spiegel/pi/cmd"
)

func main() {
	cmd.Execute()
	os.Exit(cmd.ExitCode)
}

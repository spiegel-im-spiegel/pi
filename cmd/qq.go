/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/pi/qq"
)

// qqCmd represents the qq command
var qqCmd = &cobra.Command{
	Use:   "qq [data file]",
	Short: "make Q-Q plot data",
	Long:  "make Q-Q plot data.",
	Run: func(cmd *cobra.Command, args []string) {
		inp := os.Stdin
		if len(args) > 0 {
			file, err := os.Open(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				ExitCode = 1
				return
			}
			defer file.Close()
			inp = file
		}
		if err := qq.Execute(qq.NewContext(inp, cmd.OutOrStdout(), os.Stderr)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
		}
	},
}

func init() {
	RootCmd.AddCommand(qqCmd)
}

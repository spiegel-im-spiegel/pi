/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/pi/plot"
)

// plotCmd represents the plot command
var plotCmd = &cobra.Command{
	Use:   "plot",
	Short: "Plot random points",
	Long:  "Plot random points in 0 <= x <= 1.0 and 0 <= y <= 1.0 area.",
	Run: func(cmd *cobra.Command, args []string) {
		rng, err := RngType()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
			return
		}
		if err := plot.Execute(plot.NewContext(cmd.OutOrStdout(), os.Stderr, rng, pointCount)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
		}
	},
}

func init() {
	RootCmd.AddCommand(plotCmd)
}

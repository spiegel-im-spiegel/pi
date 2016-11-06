package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/pi/plot"
)

// Flags
var (
	pointCountP int64
)

// plotCmd represents the plot command
var plotCmd = &cobra.Command{
	Use:   "plot",
	Short: "Plot random points",
	Long:  "Plot random points in 0 <= x <= 1.0 and 0 <= y <= 1.0 area.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := plot.Execute(plot.NewContext(cmd.OutOrStdout(), os.Stderr, pointCountP)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
		}
	},
}

func init() {
	RootCmd.AddCommand(plotCmd)
	plotCmd.PersistentFlags().Int64VarP(&pointCountP, "pcount", "p", 10000, "Count of points")
}

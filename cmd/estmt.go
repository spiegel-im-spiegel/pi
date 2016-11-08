package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/pi/estmt"
)

// estmtCmd represents the estmt command
var estmtCmd = &cobra.Command{
	Use:   "estmt",
	Short: "Estimate of Pi",
	Long:  "Estimate of Pi with Monte Carlo method.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := estmt.Execute(estmt.NewContext(cmd.OutOrStdout(), os.Stderr, pointCount, estmtCount)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
		}
	},
}

func init() {
	RootCmd.AddCommand(estmtCmd)
	estmtCmd.PersistentFlags().Int64VarP(&pointCount, "pcount", "p", 10000, "Count of points")
	estmtCmd.PersistentFlags().Int64VarP(&estmtCount, "ecount", "e", 100, "Count of estimate")
}

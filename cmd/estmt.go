package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/pi/estmt"
)

// Flags
var (
	pointCountE int64
	estmtCountE int64
)

// estmtCmd represents the estmt command
var estmtCmd = &cobra.Command{
	Use:   "estmt",
	Short: "Estimate of Pi",
	Long:  "Estimate of Pi with Monte Carlo method.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := estmt.Execute(estmt.NewContext(cmd.OutOrStdout(), os.Stderr, pointCountE, estmtCountE)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
		}
	},
}

func init() {
	RootCmd.AddCommand(estmtCmd)
	estmtCmd.PersistentFlags().Int64VarP(&pointCountE, "pcount", "p", 10000, "Count of points")
	estmtCmd.PersistentFlags().Int64VarP(&estmtCountE, "ecount", "e", 100, "Count of estimate")
}

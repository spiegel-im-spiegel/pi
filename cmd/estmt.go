/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */
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
		rng, err := RngType()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
			return
		}
		if err := estmt.Execute(estmt.NewContext(cmd.OutOrStdout(), os.Stderr, rng, pointCount, estmtCount)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			ExitCode = 1
		}
	},
}

func init() {
	RootCmd.AddCommand(estmtCmd)
	estmtCmd.PersistentFlags().Int64VarP(&estmtCount, "ecount", "e", 100, "Count of estimate")
}

/**
 * These codes are licensed under CC0.
 * http://creativecommons.org/publicdomain/zero/1.0/
 */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/pi/gencmplx"
)

// Flags
var (
	ExitCode   int
	rngType    string
	pointCount int64
	estmtCount int64
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pi",
	Short: "Estimate of Pi",
	Long:  "Estimate of Pi with Monte Carlo method.",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		ExitCode = 1
	}
}

func init() {
	ExitCode = 0
	RootCmd.PersistentFlags().Int64VarP(&pointCount, "pcount", "p", 10000, "Count of points")
	RootCmd.PersistentFlags().StringVarP(&rngType, "rsource", "r", "GO", "Source of RNG (GO/LCG/MT)")
}

//RngType returns kind of RNG
func RngType() (gencmplx.RNGs, error) {
	switch rngType {
	case "MT":
		return gencmplx.MT, nil
	case "LCG":
		return gencmplx.LCG, nil
	case "GO":
		return gencmplx.GO, nil
	default:
		return gencmplx.NULL, fmt.Errorf("invalid -rsource parameter: %s\n", rngType)
	}
}

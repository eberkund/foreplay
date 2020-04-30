package cmd

import (
	"foreplay/run"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [hook]",
	Short: "Run hooks.",
	Args:  cobra.NoArgs,
	Run:   runRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	r, err := run.InitializeRunner()
	if err != nil {
		panic(err)
	}
	r.Start()
}

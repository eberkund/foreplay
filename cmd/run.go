package cmd

import (
	"foreplay/run"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [hook]",
	Short: "Run hooks",
	Args:  cobra.NoArgs,
	RunE:  runRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) error {
	r, err := run.InitializeRunner()
	if err != nil {
		return err
	}
	return r.Start()
}

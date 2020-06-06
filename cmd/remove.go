package cmd

import (
	"fmt"

	"foreplay/install"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Pull out of .git/hooks",
	Args:  cobra.NoArgs,
	RunE:  runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func runRemove(cmd *cobra.Command, args []string) error {
	fmt.Println("deleting", install.PreCommitHookPath())
	return install.Remove()
}

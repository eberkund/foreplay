package cmd

import (
	"fmt"

	"foreplay/install"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Args:  cobra.NoArgs,
	Short: "Install to .git/hooks",
	RunE:  runInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func runInstall(cmd *cobra.Command, args []string) error {
	fmt.Println("installing to", install.PreCommitHookPath())
	return install.Install()
}

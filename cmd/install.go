package cmd

import (
	"foreplay/install"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Args:  cobra.NoArgs,
	Short: "Install shims into `.git/hooks.`",
	RunE:  runInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func runInstall(cmd *cobra.Command, args []string) error {
	return install.Install()
}

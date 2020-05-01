package cmd

import (
	"log"

	"foreplay/install"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Args:  cobra.NoArgs,
	Short: "Install shims into `.git/hooks.`",
	Run:   runInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func runInstall(cmd *cobra.Command, args []string) {
	if err := install.Install(); err != nil {
		log.Fatal(err)
	}
}

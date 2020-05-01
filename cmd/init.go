package cmd

import (
	"log"

	"foreplay/install"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Args:  cobra.NoArgs,
	Short: "Initializes the repo with an empty config file.",
	Run:   runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	if err := install.Init(); err != nil {
		log.Fatal(err)
	}
}

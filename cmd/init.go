package cmd

import (
	"log"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var fs = afero.NewOsFs()

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
	err := afero.WriteFile(fs, ".foreplay.yml", []byte(`hooks:
#  - id: golangci-lint
#    run: run
`), 0755)

	if err != nil {
		log.Fatal(err)
	}
}

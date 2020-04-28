package cmd

import (
	"github.com/spf13/afero"
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
	_ = afero.WriteFile(afero.OsFs{}, ".foreplay.yml", []byte(`hooks:
#  - id: golangci-lint
#    command: golangci-lint
#    args: run
#    hook: pre-commit
`), 0755)
}

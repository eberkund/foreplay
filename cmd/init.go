package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the repo with an empty config file.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runInit,
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

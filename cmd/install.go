package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

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
	fmt.Println("installing to", path.Join(".git", "hooks"))
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	preCommitHookPath := path.Join(wd, ".git", "hooks", "pre-commit")
	contents := `#!/usr/bin/env bash
exec foreplay run
`
	err = ioutil.WriteFile(preCommitHookPath, []byte(contents), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

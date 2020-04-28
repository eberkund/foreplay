package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
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
	fmt.Println("install called")
	wd, _ := os.Getwd()
	preCommitHookPath := path.Join(wd, ".git", "hooks", "pre-commit")
	_ = ioutil.WriteFile(preCommitHookPath, []byte(`#!/usr/bin/env bash
exec foreplay run
`), 0755)
}

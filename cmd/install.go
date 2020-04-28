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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runInstall,
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

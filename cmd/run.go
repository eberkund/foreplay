package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runRun,
}

type Config struct {
	Hooks []Hook
}

type Hook struct {
	Id      string
	Command string
	Args    []string
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	fmt.Println("run called")
	c := &Config{Hooks: []Hook{
		{
			Id:      "golangci-lint",
			Command: "golangci-lint",
			Args:    []string{"run"},
		},
	}}

	for _, v := range c.Hooks {
		cmd := exec.Command(v.Command, v.Args...)
		out, err := cmd.CombinedOutput()

		wd, _ := os.Getwd()
		fmt.Println(wd)

		fmt.Println(string(out))

		if err != nil {
			fmt.Println("problem encountered: ", err)
		} else {
			fmt.Println("all ok")
		}
	}
}

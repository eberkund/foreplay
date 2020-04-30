package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// stdinCmd represents the stdin command
var stdinCmd = &cobra.Command{
	Use:   "stdin",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		exe, _ := os.Executable()
		fmt.Printf("args=%+v\n", os.Args)
		fmt.Printf("exec=%s\n", exe)
		fmt.Printf("stdin=%s\n", os.Stdin.Name())
		fmt.Printf("stdout=%s\n", os.Stdout.Name())
	},
}

func init() {
	rootCmd.AddCommand(stdinCmd)
}

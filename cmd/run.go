package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [hook]",
	Short: "Run hooks.",
	Args:  cobra.MaximumNArgs(1),
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
	Id      string   `yaml:"id"`
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
	Working string   `yaml:"working"`
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
	fmt.Println("run called")
	var c Config
	data, err := ioutil.ReadFile(".foreplay.yml")
	if err != nil {
		println("could not read config file")
		panic(err)
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		println("could not parse config file")
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(c.Hooks))

	wd, _ := os.Getwd()
	fmt.Println(wd)

	fmt.Printf("%d hook(s) listed\n", len(c.Hooks))

	for _, v := range c.Hooks {

		go func(v Hook) {
			println(v.Id)

			if v.Working != "" {
				err := os.Chdir(v.Working)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}

			cmd := exec.Command(v.Command, v.Args...)

			out, err := cmd.CombinedOutput()
			fmt.Println(string(out))

			if err != nil {
				fmt.Println("problem encountered: ", err)
				os.Exit(1)
			}
			wg.Done()
		}(v)

	}

	wg.Wait()
}

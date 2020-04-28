package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
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
	Hooks []*Hook
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

	wd, _ := os.Getwd()
	fmt.Println(wd)

	fmt.Printf("%d hook(s) listed\n", len(c.Hooks))
	group := errgroup.Group{}

	for _, hook := range c.Hooks {
		hook := hook
		group.Go(func() error {
			println(hook.Id)

			p := path.Join(wd, hook.Working)
			println(p)
			err := os.Chdir(p)
			if err != nil {
				return err
			}

			out, err := exec.Command(hook.Command, hook.Args...).CombinedOutput()
			if err != nil {
				return err
			}

			fmt.Println(string(out))
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		os.Exit(1)
	}
}

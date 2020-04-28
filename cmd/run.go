package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/pkg/errors"
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
	Hooks []Hook `yaml:"hooks" jsonschema:"required"`
}

type Hook struct {
	ID  string `yaml:"id" jsonschema:"required"`
	Run string `yaml:"run" jsonschema:"required"`
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRun(cmd *cobra.Command, args []string) {
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

	//
	//fmt.Printf("%d hook(s) listed\n", len(c.Hooks))
	group := errgroup.Group{}

	s := spinner.New(spinner.CharSets[11], 125*time.Millisecond)

	for _, hook := range c.Hooks {
		hook := hook
		println(hook.ID)
		//spinners.AddSpinner(hook.ID)
		group.Go(func() error {
			cmd := exec.Command("sh")
			cmd.Stdin = bytes.NewBuffer([]byte(hook.Run))
			_, err := cmd.CombinedOutput()

			if err != nil {
				//fmt.Println(string(out))
				return errors.Wrapf(err, "error running %q", hook.ID)
			}

			return nil
		})
	}

	s.Start()
	err = group.Wait()
	s.Stop()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

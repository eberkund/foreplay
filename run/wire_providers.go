package run

import (
	"errors"
	"os"
	"os/exec"

	"foreplay/config"
	"foreplay/output"
	"foreplay/output/common"
)

type Shell = string

func GetConfig() (config.Config, error) {
	return config.Get()
}

func GetShell() (Shell, error) {
	shells := []string{
		"sh",
		"bash",
		"powershell",
	}
	for _, shell := range shells {
		cmd, err := exec.LookPath(shell)
		if err == nil {
			return cmd, nil
		}
	}
	return "", errors.New("could not find shell to execute in")
}

func GetPrinter(c config.Config) common.Registerable {
	return output.GetOutput(c.Style, os.Stdout)
}

func GetRun(
	shell Shell,
	c config.Config,
	printer common.Registerable,
) *Run {
	return &Run{
		Shell:   shell,
		Printer: printer,
		Hooks:   c.Hooks,
		Timeout: c.Timeout,
	}
}

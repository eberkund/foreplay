package run

import (
	"os"
	"runtime"

	"foreplay/config"
	"foreplay/output"
	"foreplay/output/common"
)

func GetConfig() (config.Config, error) {
	return config.Get()
}

func GetShell() string {
	var command string
	if runtime.GOOS == "windows" {
		command = "C:/Program Files/Git/usr/bin/sh.exe"
	} else {
		command = "sh"
	}
	return command
}

func GetPrinter(c config.Config) common.Registerable {
	return output.GetOutput(c.Style, os.Stdout)
}

func GetExit() func(int) {
	return os.Exit
}

func GetRun(
	command string,
	c config.Config,
	printer common.Registerable,
	exit func(int),
) *Run {
	return &Run{
		Shell:   command,
		Printer: printer,
		Hooks:   c.Hooks,
		exit:    exit,
	}
}

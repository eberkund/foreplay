package output

import (
	"os"

	"foreplay/output/common"
	"foreplay/output/plain"
	"foreplay/output/spinner"

	"github.com/k0kubun/go-ansi"
)

func GetOutput(output string) common.Registerable {
	switch output {
	case "plain":
		return plain.New(os.Stdout)
	case "spinner":
		return spinner.New(ansi.NewAnsiStdout())
	}
	return nil
}

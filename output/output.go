package output

import (
	"io"

	"foreplay/output/common"
	"foreplay/output/plain"
	"foreplay/output/spinner"

	"github.com/shiena/ansicolor"
)

func GetOutput(output string, writer io.Writer) common.Registerable {
	switch output {
	case "plain":
		return plain.New(writer)
	case "spinner":
		return spinner.New(ansicolor.NewAnsiColorWriter(writer))
	}
	return nil
}

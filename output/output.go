package output

import (
	"io"

	"github.com/eberkund/foreplay/output/common"
	"github.com/eberkund/foreplay/output/plain"
	"github.com/eberkund/foreplay/output/spinner"

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

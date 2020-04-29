package output

import (
	"foreplay/output/common"
	"foreplay/output/plain"
	"foreplay/output/spinner"
)

func GetOutput(output string) common.Registerable {
	switch output {
	case "plain":
		return plain.New()
	case "spinner":
		return spinner.New()
	}
	return nil
}

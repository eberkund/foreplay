package output

import (
	"testing"

	"foreplay/output/common"
	"foreplay/output/plain"
	"foreplay/output/spinner"

	"github.com/stretchr/testify/assert"
)

func TestParseDriver(t *testing.T) {
	cases := []struct {
		Input    string
		Expected common.Registerable
	}{
		{"spinner", spinner.New(nil)},
		{"plain", plain.New(nil)},
		{"foobar", nil},
	}
	for _, tc := range cases {
		t.Run(tc.Input, func(t *testing.T) {
			output := GetOutput(tc.Input)
			assert.IsType(t, tc.Expected, output)
		})
	}
}

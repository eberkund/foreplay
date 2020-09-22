package output

import (
	"os"
	"testing"

	"github.com/eberkund/foreplay/output/common"
	"github.com/eberkund/foreplay/output/plain"
	"github.com/eberkund/foreplay/output/spinner"

	"github.com/stretchr/testify/require"
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
			output := GetOutput(tc.Input, os.Stdout)
			require.IsType(t, tc.Expected, output)
		})
	}
}

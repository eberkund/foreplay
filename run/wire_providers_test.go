package run

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetShell(t *testing.T) {
	err := os.Setenv("PATH", "")
	require.NoError(t, err)

	_, err = GetShell()
	require.Errorf(t, err, "could not find shell to execute in")
}

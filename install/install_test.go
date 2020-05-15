package install

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestInstall(t *testing.T) {
	Fs = afero.NewMemMapFs()
	err := Install()

	require.NoError(t, err)
	contents, err := afero.ReadFile(Fs, PreCommitHookPath())
	require.NoError(t, err)
	require.Equal(t, `#!/bin/sh
exec &> /dev/tty
foreplay run
`, string(contents))
}

func TestInit(t *testing.T) {
	Fs = afero.NewMemMapFs()

	err := Init()

	require.NoError(t, err)
	contents, err := afero.ReadFile(Fs, ".foreplay.yml")
	require.NoError(t, err)
	require.Equal(t, `hooks:
#  - id: golangci-lint
#    run: run
`, string(contents))
}

package install

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestInstall(t *testing.T) {
	fs = afero.NewMemMapFs()
	err := Install()

	require.NoError(t, err)
	contents, err := afero.ReadFile(fs, PreCommitHookPath())
	require.NoError(t, err)
	require.Equal(t, `#!/usr/bin/env bash
exec foreplay run
`, string(contents))
}

func TestInit(t *testing.T) {
	fs = afero.NewMemMapFs()

	err := Init()

	require.NoError(t, err)
	contents, err := afero.ReadFile(fs, ".foreplay.yml")
	require.NoError(t, err)
	require.Equal(t, `hooks:
#  - id: golangci-lint
#    run: run
`, string(contents))
}

package cmd

import (
	"bytes"
	"testing"

	"foreplay/config"
	"foreplay/install"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestRootCommand(t *testing.T) {
	var b bytes.Buffer
	rootCmd.SetOut(&b)
	rootCmd.SetArgs([]string{})

	Execute()
	require.Contains(t, b.String(), "foreplay [command]")
}

func TestInitCommand(t *testing.T) {
	install.Fs = afero.NewMemMapFs()

	rootCmd.SetArgs([]string{"init"})
	err := rootCmd.Execute()
	require.NoError(t, err)
}

func TestInstallCommand(t *testing.T) {
	install.Fs = afero.NewMemMapFs()

	rootCmd.SetArgs([]string{"install"})
	err := rootCmd.Execute()
	require.NoError(t, err)
}

func TestSchemaCommand(t *testing.T) {
	var b bytes.Buffer
	rootCmd.SetOut(&b)
	rootCmd.SetArgs([]string{"schema"})

	expected, err := config.Schema()
	require.NoError(t, err)

	err = rootCmd.Execute()
	require.NoError(t, err)
	require.Equal(t, string(expected), b.String())
}

func TestRunCommandWithNoConfig(t *testing.T) {
	config.Fs = afero.NewMemMapFs()

	var b bytes.Buffer
	rootCmd.SetOut(&b)
	rootCmd.SetArgs([]string{"run"})

	err := rootCmd.Execute()
	require.Error(t, err)
}

func TestRunCommand(t *testing.T) {
	config.Fs = afero.NewMemMapFs()
	err := afero.WriteFile(config.Fs, ".foreplay.yml", []byte(`
style: plain
hooks:
  - id: pwd
    run: pwd
`), 0755)
	require.NoError(t, err)

	var b bytes.Buffer
	rootCmd.SetOut(&b)
	rootCmd.SetArgs([]string{"run"})

	err = rootCmd.Execute()
	require.NoError(t, err)
}

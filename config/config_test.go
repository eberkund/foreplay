package config

import (
	"io/ioutil"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoConfigFile(t *testing.T) {
	fs = afero.NewMemMapFs()
	_, err := Get()
	assert.Error(t, err)
}

func TestCanReadConfigFile(t *testing.T) {
	cfg, err := ioutil.ReadFile("testdata/good-config.yml")
	require.NoError(t, err)

	fs = afero.NewMemMapFs()
	err = afero.WriteFile(fs, ".foreplay.yml", cfg, 0755)
	require.NoError(t, err)

	c, err := Get()
	require.NoError(t, err)
	require.NotNil(t, c)
}

func TestBadConfigFile(t *testing.T) {
	cfg, err := ioutil.ReadFile("testdata/bad-config.yml")
	require.NoError(t, err)

	fs = afero.NewMemMapFs()
	err = afero.WriteFile(fs, ".foreplay.yml", cfg, 0755)
	require.NoError(t, err)

	_, err = Get()
	require.Error(t, err)
}

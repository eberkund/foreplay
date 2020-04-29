package config

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNoConfigFile(t *testing.T) {
	fs = afero.NewMemMapFs()
	_, err := Get()
	assert.Error(t, err)
}

//func TestCanReadConfigFile(t *testing.T) {
//	fs = afero.NewMemMapFs()
//	_, err := Get()
//	assert.Error(t, err)
//}

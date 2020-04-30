package config

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaMatches(t *testing.T) {
	pre, err := ioutil.ReadFile(path.Join("..", "schema.json"))
	require.NoError(t, err)

	generated, err := Schema()
	require.NoError(t, err)

	require.Equal(
		t,
		string(generated),
		string(pre),
	)
}

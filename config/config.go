package config

import (
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

var Fs = afero.NewOsFs()

const configFile = ".foreplay.yml"

// Config represents the `.foreplay.yml` schema.
type Config struct {
	Style string `yaml:"style"`
	Hooks []Hook `yaml:"hooks" jsonschema:"required"`
}

// Hook represents a task to be run on pre-commit.
type Hook struct {
	ID  string `yaml:"id" jsonschema:"required"`
	Run string `yaml:"run" jsonschema:"required"`
}

// Parses the config file and returns a struct.
func Get() (Config, error) {
	var c Config
	data, err := afero.ReadFile(Fs, configFile)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

package config

import (
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

var fs = afero.NewOsFs()

const configFile = ".foreplay.yml"

type Config struct {
	Hooks []Hook `yaml:"hooks" jsonschema:"required"`
}

type Hook struct {
	ID  string `yaml:"id" jsonschema:"required"`
	Run string `yaml:"run" jsonschema:"required"`
}

func Get() (*Config, error) {
	var c Config
	data, err := afero.ReadFile(fs, configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

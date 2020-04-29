package config

import (
	"encoding/json"

	"github.com/alecthomas/jsonschema"
	"honnef.co/go/tools/config"
)

func Schema() ([]byte, error) {
	reflector := jsonschema.Reflector{}
	schema := reflector.Reflect(&config.Config{})
	return json.MarshalIndent(schema, "", "  ")
}

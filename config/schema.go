package config

import (
	"encoding/json"

	"github.com/alecthomas/jsonschema"
)

func Schema() ([]byte, error) {
	reflector := jsonschema.Reflector{}
	schema := reflector.Reflect(&Config{})
	return json.MarshalIndent(schema, "", "  ")
}

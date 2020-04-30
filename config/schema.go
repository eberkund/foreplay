package config

import (
	"encoding/json"

	"github.com/alecthomas/jsonschema"
)

// Schema returns the config file JSON schema.
func Schema() ([]byte, error) {
	reflector := jsonschema.Reflector{}
	schema := reflector.Reflect(&Config{})
	return json.MarshalIndent(schema, "", "  ")
}

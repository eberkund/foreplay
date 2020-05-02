package config

import (
	"encoding/json"

	"github.com/alecthomas/jsonschema"
)

// Schema returns the config file JSON schema.
func Schema() []byte {
	reflector := jsonschema.Reflector{}
	schema := reflector.Reflect(&Config{})
	jsonSchema, _ := json.MarshalIndent(schema, "", "  ")
	return jsonSchema
}

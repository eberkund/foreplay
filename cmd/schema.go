package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/jsonschema"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Display config JSON schema",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		printConfigSchema()
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}

func printConfigSchema() {
	schema := jsonschema.Reflect(&Config{})
	result, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		logrus.WithError(err).Fatal("could not marshal schema to JSON")
	}
	fmt.Print(string(result))
}

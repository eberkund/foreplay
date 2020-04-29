package cmd

import (
	"fmt"

	"foreplay/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Display the config file JSON schema.",
	Args:  cobra.NoArgs,
	Run:   printConfigSchema,
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}

func printConfigSchema(cmd *cobra.Command, args []string) {
	schema, err := config.Schema()
	if err != nil {
		logrus.WithError(err).Fatal("could not marshal schema to JSON")
	}
	fmt.Print(string(schema))
}

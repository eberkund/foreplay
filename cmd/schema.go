package cmd

import (
	"fmt"

	"foreplay/config"

	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:    "schema",
	Short:  "Display the config file JSON schema",
	Args:   cobra.NoArgs,
	RunE:   printConfigSchema,
	Hidden: true,
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}

func printConfigSchema(cmd *cobra.Command, args []string) error {
	schema := config.Schema()
	_, err := fmt.Fprint(cmd.OutOrStderr(), string(schema))
	return err
}

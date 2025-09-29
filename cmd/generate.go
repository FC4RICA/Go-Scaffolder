package cmd

import (
	"fmt"

	"github.com/FC4RICA/Go-Scaffolder/internal/generator/parser"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate [domain file]",
	Short: "generate go files from struct in domain package",
	RunE: func(cmd *cobra.Command, args []string) error {
		domainFile := args[0]

		entity, err := parser.ParseDomain(domainFile)
		if err != nil {
			return err
		}

		fmt.Println(entity)

		return nil
	},
}

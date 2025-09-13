/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/FC4RICA/Go-Scaffolder/internal/scaffold"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create an empty project",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		if err := scaffold.CreateProject(projectName); err != nil {
			return fmt.Errorf("failed to create project %s: %w", projectName, err)
		}

		fmt.Printf("project %s created successfully\n", projectName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

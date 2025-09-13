package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-scaffolder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go Scaffolder Generator v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

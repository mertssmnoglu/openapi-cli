package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openapi-cli",
	Short: "A CLI for OpenAPI operations",
	Long:  `A Command Line Interface (CLI) tool for serving, validating, and linting OpenAPI specifications.`,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

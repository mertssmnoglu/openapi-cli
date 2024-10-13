package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/flowchartsman/swaggerui"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the OpenAPI file",
	Long:  `Starts a server to serve the OpenAPI file.`,
	Run:   CommandServe,
}

func init() {
	serveCmd.Flags().String("file", "./openapi.json", "The OpenAPI file to serve")
}

func ReadOpenAPIFile(filePath string) ([]byte, error) {
	// Read the OpenAPI file
	// OpenAPI Specification files are in JSON or YAML format
	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file:", err)

		return nil, err
	}

	return data, nil
}

func CommandServe(cmd *cobra.Command, args []string) {
	filePath, err := cmd.Flags().GetString("file")

	if err != nil {
		fmt.Println("Error reading file flag:", err)
		return
	}

	content, err := ReadOpenAPIFile(filePath)
	fmt.Println(string(content))

	http.Handle("/openapi/", http.StripPrefix("/openapi", swaggerui.Handler(content)))
	log.Println("serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
